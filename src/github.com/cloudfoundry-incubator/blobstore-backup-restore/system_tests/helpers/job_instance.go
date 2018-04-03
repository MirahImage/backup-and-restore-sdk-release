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

package helpers

import (
	"encoding/json"
	"io"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"strings"
)

type JobInstance struct {
	Deployment    string
	Instance      string
	InstanceIndex string
}

type jsonOutputFromCli struct {
	Tables []struct {
		Rows []map[string]string
	}
}

func RunCommand(cmd string, args ...string) *gexec.Session {
	return runCommandWithStream(GinkgoWriter, GinkgoWriter, cmd, args...)
}

func runCommandWithStream(stdout, stderr io.Writer, cmd string, args ...string) *gexec.Session {
	cmdParts := strings.Split(cmd, " ")
	commandPath := cmdParts[0]
	combinedArgs := append(cmdParts[1:], args...)
	command := exec.Command(commandPath, combinedArgs...)

	session, err := gexec.Start(command, stdout, stderr)

	Expect(err).ToNot(HaveOccurred())
	Eventually(session).Should(gexec.Exit())
	return session
}

func (jobInstance *JobInstance) RunOnVMAndSucceed(command string) *gexec.Session {
	session := jobInstance.RunOnInstance(command)
	Expect(session).To(gexec.Exit(0), string(session.Err.Contents()))

	return session
}

func (jobInstance *JobInstance) RunOnInstance(cmd ...string) *gexec.Session {
	return RunCommand(
		join(
			BoshCommand(),
			forDeployment(jobInstance.Deployment),
			getSSHCommand(jobInstance.Instance, jobInstance.InstanceIndex),
		),
		join(cmd...),
	)
}

func (jobInstance *JobInstance) getIPOfInstance() string {
	session := RunCommand(
		BoshCommand(),
		forDeployment(jobInstance.Deployment),
		"instances",
		"--json",
	)
	outputFromCli := jsonOutputFromCli{}
	contents := session.Out.Contents()
	Expect(json.Unmarshal(contents, &outputFromCli)).To(Succeed())
	for _, instanceData := range outputFromCli.Tables[0].Rows {
		if strings.HasPrefix(instanceData["Instance"], jobInstance.Instance+"/") {
			return instanceData["ips"]
		}
	}
	Fail("Cant find instances with name '" + jobInstance.Instance + "' and Deployment name '" + jobInstance.Deployment + "'")
	return ""
}

func (jobInstance *JobInstance) DownloadFromInstance(remotePath, localPath string) *gexec.Session {
	return RunCommand(
		join(
			BoshCommand(),
			forDeployment(jobInstance.Deployment),
			getDownloadCommand(remotePath, localPath, jobInstance.Instance, jobInstance.InstanceIndex),
		),
	)
}

func (jobInstance *JobInstance) UploadToInstance(localPath, remotePath string) *gexec.Session {
	return RunCommand(
		join(
			BoshCommand(),
			forDeployment(jobInstance.Deployment),
			getUploadCommand(localPath, remotePath, jobInstance.Instance, jobInstance.InstanceIndex),
		),
	)
}

func join(args ...string) string {
	return strings.Join(args, " ")
}
