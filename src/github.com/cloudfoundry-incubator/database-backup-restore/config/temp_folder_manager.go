package config

import (
	"io/ioutil"
	"os"
)

type TempFolderManager struct {
	FolderPath string
}

func NewTempFolderManager() (TempFolderManager, error) {
	folderPath, err := ioutil.TempDir("", "")
	if err != nil {
		return TempFolderManager{}, err
	}
	return TempFolderManager{FolderPath: folderPath}, nil
}

func (m TempFolderManager) WriteTempFile(contents string) (string, error) {
	file, err := ioutil.TempFile(m.FolderPath, "")
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(file.Name(), []byte(contents), 0777)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func (m TempFolderManager) Cleanup() {
	os.RemoveAll(m.FolderPath)
}
