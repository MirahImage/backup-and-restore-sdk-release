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
	"strconv"

	. "github.com/onsi/ginkgo"

	"fmt"
	"time"
)

var _ = Describe("postgres-backup", func() {
	Context("database-backup-restorer is collocated with Postgres", func() {
		var dbDumpPath string
		var configPath string
		var databaseName string
		var dbJob JobInstance
		postgresPackage := "postgres-9.6.3"

		BeforeEach(func() {
			dbJob = JobInstance{
				deployment:    "postgres-9.6-dev",
				instance:      "postgres",
				instanceIndex: "0",
			}

			configPath = "/tmp/config.json" + strconv.FormatInt(time.Now().Unix(), 10)
			dbDumpPath = "/tmp/sql_dump" + strconv.FormatInt(time.Now().Unix(), 10)
			databaseName = "db" + strconv.FormatInt(time.Now().Unix(), 10)

			dbJob.runOnVMAndSucceed(
				fmt.Sprintf(`/var/vcap/packages/postgres-9.6.3/bin/createdb -U vcap "%s"`, databaseName))
			dbJob.runPostgresSqlCommand("CREATE TABLE people (name varchar);", databaseName, postgresPackage)
			dbJob.runPostgresSqlCommand("INSERT INTO people VALUES ('Derik');", databaseName, postgresPackage)

			configJson := fmt.Sprintf(
				`{"username":"vcap","password":"%s","host":"localhost","port":5432,"database":"%s","adapter":"postgres"}`,
				MustHaveEnv("POSTGRES_PASSWORD"),
				databaseName,
			)

			dbJob.runOnVMAndSucceed(fmt.Sprintf("echo '%s' > %s", configJson, configPath))
		})

		AfterEach(func() {
			dbJob.runOnVMAndSucceed(fmt.Sprintf(`/var/vcap/packages/postgres-9.6.3/bin/dropdb -U vcap "%s"`, databaseName))
			dbJob.runOnVMAndSucceed(fmt.Sprintf("rm -rf %s %s", configPath, dbDumpPath))
		})

		It("backs up the Postgres database", func() {
			dbJob.runOnVMAndSucceed(
				fmt.Sprintf(
					"/var/vcap/jobs/database-backup-restorer/bin/backup --artifact-file %s --config %s",
					dbDumpPath,
					configPath,
				),
			)
			dbJob.runOnVMAndSucceed(fmt.Sprintf("ls -l %s", dbDumpPath))
		})
	})
})