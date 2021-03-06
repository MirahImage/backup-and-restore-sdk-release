package blobstore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//go:generate counterfeiter -o fakes/fake_versioned_artifact.go . VersionedArtifact
type VersionedArtifact interface {
	Save(backup map[string]BucketSnapshot) error
	Load() (map[string]BucketSnapshot, error)
}

type BucketSnapshot struct {
	BucketName string        `json:"bucket_name"`
	RegionName string        `json:"region_name"`
	Versions   []BlobVersion `json:"versions"`
}

type BlobVersion struct {
	BlobKey string `json:"blob_key"`
	Id      string `json:"version_id"`
}

type VersionedFileArtifact struct {
	filePath string
}

func NewVersionedFileArtifact(filePath string) VersionedFileArtifact {
	return VersionedFileArtifact{filePath: filePath}
}

func (a VersionedFileArtifact) Save(backup map[string]BucketSnapshot) error {
	marshalledBackup, err := json.MarshalIndent(backup, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(a.filePath, marshalledBackup, 0666)
	if err != nil {
		return fmt.Errorf("could not write backup file: %s", err.Error())
	}

	return nil
}

func (a VersionedFileArtifact) Load() (map[string]BucketSnapshot, error) {
	bytes, err := ioutil.ReadFile(a.filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read backup file: %s", err.Error())
	}

	var backup map[string]BucketSnapshot
	err = json.Unmarshal(bytes, &backup)
	if err != nil {
		return nil, fmt.Errorf("backup file has an invalid format: %s", err.Error())
	}

	return backup, nil
}
