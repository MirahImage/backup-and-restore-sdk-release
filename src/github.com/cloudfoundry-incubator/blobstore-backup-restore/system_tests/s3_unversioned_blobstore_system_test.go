// Copyright (C) 2017-Present Pivotal Software, Inc. All rights reserved.
//
// This program and the accompanying materials are made available under
// the terms of the under the Apache License, Version 2.0 (the "License”);
// you may not use this file except in compliance with the License.
//
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.

package system_tests

import (
	"time"

	"strconv"

	"io/ioutil"

	"os"

	. "github.com/cloudfoundry-incubator/blobstore-backup-restore/system_tests/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("S3 unversioned backuper", func() {
	var region string
	var bucket string
	var backupRegion string
	var backupBucket string
	var instanceArtifactDirPath string

	var blobKey string
	var localArtifact *os.File
	var backuperInstance JobInstance

	BeforeEach(func() {
		backuperInstance = JobInstance{
			Deployment:    MustHaveEnv("BOSH_DEPLOYMENT"),
			Instance:      "s3-unversioned-backuper",
			InstanceIndex: "0",
		}

		region = MustHaveEnv("S3_UNVERSIONED_BUCKET_REGION")
		bucket = MustHaveEnv("S3_UNVERSIONED_BUCKET_NAME")

		backupRegion = MustHaveEnv("S3_UNVERSIONED_BACKUP_BUCKET_REGION")
		backupBucket = MustHaveEnv("S3_UNVERSIONED_BACKUP_BUCKET_NAME")

		DeleteAllFilesFromBucket(region, bucket)
		DeleteAllFilesFromBucket(backupRegion, backupBucket)

		instanceArtifactDirPath = "/tmp/s3-unversioned-blobstore-backup-restorer" + strconv.FormatInt(time.Now().Unix(), 10)
		backuperInstance.RunOnVMAndSucceed("mkdir -p " + instanceArtifactDirPath)
		var err error
		localArtifact, err = ioutil.TempFile("", "blobstore-")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		backuperInstance.RunOnVMAndSucceed("rm -rf " + instanceArtifactDirPath)
		err := os.Remove(localArtifact.Name())
		Expect(err).NotTo(HaveOccurred())
		DeleteAllFilesFromBucket(region, bucket)
		DeleteAllFilesFromBucket(backupRegion, backupBucket)
	})

	It("backs up from the source bucket to the backup bucket", func() {
		blobKey = UploadTimestampedFileToBucket(region, bucket, "some/folder/file1", "FILE1")

		backuperInstance.RunOnVMAndSucceed("BBR_ARTIFACT_DIRECTORY=" + instanceArtifactDirPath +
			" /var/vcap/jobs/s3-unversioned-blobstore-backup-restorer/bin/bbr/backup")

		filesList := ListFilesFromBucket(backupRegion, backupBucket)

		Expect(filesList).To(ConsistOf(MatchRegexp(
			"\\d{4}_\\d{2}_\\d{2}_\\d{2}_\\d{2}_\\d{2}/my_bucket/" + blobKey + "$")))

		Expect(GetFileContentsFromBucket(backupRegion, backupBucket, filesList[0])).To(Equal("FILE1"))

		session := backuperInstance.DownloadFromInstance(
			instanceArtifactDirPath+"/blobstore.json", localArtifact.Name())
		Expect(session).Should(gexec.Exit(0))
		fileContents, err := ioutil.ReadFile(localArtifact.Name())
		Expect(err).NotTo(HaveOccurred())
		Expect(fileContents).To(ContainSubstring("\"my_bucket\": {"))
		Expect(fileContents).To(ContainSubstring("\"bucket_name\": \"" + backupBucket + "\""))
		Expect(fileContents).To(ContainSubstring("\"bucket_region\": \"" + backupRegion + "\""))
		Expect(fileContents).To(MatchRegexp(
			"\"path\": \"\\d{4}_\\d{2}_\\d{2}_\\d{2}_\\d{2}_\\d{2}\\/my_bucket\""))
	})

	PIt("connects with a blobstore with custom CA cert", func() {

	})
})
