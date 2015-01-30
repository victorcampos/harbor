package main

import (
	"flag"
	"fmt"
	"github.com/victorcampos/harbor/commandline"
	"github.com/victorcampos/harbor/config"
	"github.com/victorcampos/harbor/download"
	"github.com/victorcampos/harbor/execute"
	"github.com/victorcampos/harbor/execute/docker"
	"os"
)

func main() {
	setCustomUsageMessage()

	configVars := make(commandline.ConfigVarsMap)

	flag.Var(&configVars, "e", "sets configuration variables")
	flag.Parse()

	harborConfig, err := config.Load(configVars)
	checkError(err)

	err = download.FromS3(harborConfig)
	checkError(err)

	err = execute.Commands(harborConfig)
	checkError(err)

	err = docker.Build(harborConfig.ImageTag)
	checkError(err)
}

func setCustomUsageMessage() {
	// TODO: Alter the help message when not using /bin/bash anymore
	flag.Usage = func() {
		helpText := `Usage: harbor [-e KEY=VALUE]

Harbor looks up a file named harbor.yml in the same directory where run from, harbor.yml structure is:
 imagetag: <tag to be used on 'docker build'>
 downloadpath: <local root path to download files into>
 s3:
   bucket: <base bucket to download files from>
   basepath: <inside the bucket the root path for files to be downloaded>
 files:
   - s3path: <path to file in S3 after [s3.bucket]/[s3.basepath]>
     localname: <local path + name of the file, will be downloaded into [downloadpath]/[localname]>
     permission: <[optional] file permissions, default 0644>
 - commands:
   <YAML array containing shell commands (currently /bin/bash) to be run before 'docker build'>

 You can use ${<KEY>} as a placeholder in harbor.yml to be replaced by the value passed in a -e flag

Options:
  -e []: List of KEY=VALUE to be replaced in harbor.yml
		`
		fmt.Println(helpText)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
