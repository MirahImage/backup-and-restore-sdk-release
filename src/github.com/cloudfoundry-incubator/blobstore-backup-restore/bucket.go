package blobstore

import (
	"fmt"
)

//go:generate counterfeiter -o fakes/fake_bucket.go . Bucket
type Bucket interface {
	Name() string
	RegionName() string
	Versions() ([]Version, error)
	CopyVersions(regionName, bucketName string, versions []BlobVersion) error
}

type S3Bucket struct {
	awsCliPath string
	name       string
	regionName string
	accessKey  S3AccessKey
	endpoint   string
}

type S3AccessKey struct {
	Id     string
	Secret string
}

func NewS3Bucket(awsCliPath, name, region, endpoint string, accessKey S3AccessKey) S3Bucket {
	return S3Bucket{
		awsCliPath: awsCliPath,
		name:       name,
		regionName: region,
		accessKey:  accessKey,
		endpoint:   endpoint,
	}
}

func (bucket S3Bucket) Name() string {
	return bucket.name
}

func (bucket S3Bucket) RegionName() string {
	return bucket.regionName
}

func (bucket S3Bucket) Versions() ([]Version, error) {
	s3Api, err := NewS3API(bucket.endpoint, bucket.regionName, bucket.accessKey.Id, bucket.accessKey.Secret)

	bucketIsVersioned, err := s3Api.IsVersioned(bucket.name)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve versions from bucket %s: %s", bucket.name, err)
	}

	if !bucketIsVersioned {
		return nil, fmt.Errorf("bucket %s is not versioned", bucket.name)
	}

	versions, err := s3Api.ListObjectVersions(bucket.name)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve versions from bucket %s: %s", bucket.name, err)
	}

	return versions, nil
}

func (bucket S3Bucket) CopyVersions(sourceBucketRegion, sourceBucketName string, versionsToCopy []BlobVersion) error {
	s3Api, err := NewS3API(bucket.endpoint, bucket.regionName, bucket.accessKey.Id, bucket.accessKey.Secret)

	bucketIsVersioned, err := s3Api.IsVersioned(bucket.name)
	if err != nil {
		return err
	}

	if !bucketIsVersioned {
		return fmt.Errorf("bucket %s is not versioned", bucket.name)
	}

	if err != nil {
		return err
	}

	for _, versionToCopy := range versionsToCopy {
		err = s3Api.CopyVersion(sourceBucketName, versionToCopy.BlobKey, versionToCopy.Id, bucket.name)
		if err != nil {
			return err
		}
	}

	var keysThatShouldBePresentInBucket []string
	for _, versionToCopy := range versionsToCopy {
		keysThatShouldBePresentInBucket = append(keysThatShouldBePresentInBucket, versionToCopy.BlobKey)
	}

	return nil
}

type Version struct {
	Key      string
	Id       string `json:"VersionId"`
	IsLatest bool
}
