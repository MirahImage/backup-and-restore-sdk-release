package blobstore

import "fmt"

type Backuper struct {
	sourceBuckets       map[string]Bucket
	destinationArtifact Artifact
}

func NewBackuper(sourceBuckets map[string]Bucket, destinationArtifact Artifact) Backuper {
	return Backuper{
		sourceBuckets:       sourceBuckets,
		destinationArtifact: destinationArtifact,
	}
}

func (b Backuper) Backup() error {
	bucketSnapshots := map[string]BucketSnapshot{}

	for identifier, bucketToBackup := range b.sourceBuckets {
		versions, err := bucketToBackup.Versions()
		if err != nil {
			return err
		}

		latestVersions := filterLatest(versions)
		if containsNullVersion(latestVersions) {
			return fmt.Errorf("failed to retrieve versions; bucket '%s' has `null` VerionIds", bucketToBackup.Name())
		}

		bucketSnapshots[identifier] = BucketSnapshot{
			BucketName: bucketToBackup.Name(),
			RegionName: bucketToBackup.RegionName(),
			Versions:   latestVersions,
		}
	}

	return b.destinationArtifact.Save(bucketSnapshots)
}

func containsNullVersion(latestVersions []BlobVersion) bool {
	for _, version := range latestVersions {
		if version.Id == "null" {
			return true
		}
	}
	return false
}

func filterLatest(versions []Version) []BlobVersion {
	filteredVersions := []BlobVersion{}
	for _, version := range versions {
		if version.IsLatest {
			filteredVersions = append(filteredVersions, BlobVersion{Id: version.Id, BlobKey: version.Key})
		}
	}
	return filteredVersions
}
