package config_test

import (
	. "github.com/cloudfoundry-incubator/database-backup-restore/config"

	"os"

	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TempFolderManager", func() {
	var tempFolderManager TempFolderManager
	var err error

	BeforeEach(func() {
		tempFolderManager, err = NewTempFolderManager()
		Expect(err).NotTo(HaveOccurred())
		Expect(tempFolderManager.FolderPath).To(HavePrefix(os.TempDir()))
	})

	AfterEach(func() {
		tempFolderManager.Cleanup()
	})

	Describe("WriteTempFile", func() {
		It("creates a file in the temp folder", func() {
			filePath, err := tempFolderManager.WriteTempFile("test contents")
			Expect(err).NotTo(HaveOccurred())
			Expect(filePath).To(HavePrefix(tempFolderManager.FolderPath))
			Expect(filePath).To(BeAnExistingFile())
			Expect(ioutil.ReadFile(filePath)).To(Equal([]byte("test contents")))
		})
	})

	Describe("Cleanup", func() {
		It("removes the temp folder", func() {
			filePath, err := tempFolderManager.WriteTempFile("test contents")
			Expect(err).NotTo(HaveOccurred())
			Expect(filePath).To(HavePrefix(tempFolderManager.FolderPath))
			Expect(filePath).To(BeAnExistingFile())
			Expect(ioutil.ReadFile(filePath)).To(Equal([]byte("test contents")))

			tempFolderManager.Cleanup()
			Expect(filePath).NotTo(BeAnExistingFile())
		})
	})
})
