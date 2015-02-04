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

const VERSION = "0.1.0"

func main() {
	setCustomUsageMessage()

	configVars := make(commandline.ConfigVarsMap)

	flag.Var(&configVars, "e", "sets configuration variables")
	noDownloadFlag := flag.Bool("no-download", false, "do not download files")
	noCommandFlag := flag.Bool("no-command", false, "do not run commands")
	noDockerFlag := flag.Bool("no-docker", false, "do not run docker build, tag and push after")
	showVersionFlag := flag.Bool("v", false, "shows version")
	flag.Parse()

	if *showVersionFlag {
		fmt.Printf("Harbor version %s\n", VERSION)
		os.Exit(0)
	}

	harborConfig, err := config.Load(configVars)
	checkError(err)

	if !*noDownloadFlag {
		err = download.FromS3(harborConfig)
		checkError(err)
	}

	if !*noCommandFlag {
		err = execute.Commands(harborConfig)
		checkError(err)
	}

	if !*noDockerFlag {
		err = docker.Build(harborConfig.ImageTag, harborConfig.Tags)
		checkError(err)
	}
}

func setCustomUsageMessage() {
	// TODO: Alter the help message when not using /bin/bash anymore
	flag.Usage = func() {
		helpText := `Usage: harbor [-h] [-v] [-e <KEY>=<VALUE>] [-no-download] [-no-command] [-no-docker]

Harbor looks up a file named harbor.yml in the same directory where run from, harbor.yml structure is:
 imagetag: <tag to be used on 'docker build'>
 downloadpath: <local root path to download files into>
 s3:
   bucket: <base bucket to download files from>
   basepath: <inside the bucket the root path for files to be downloaded>
 files:
   - s3path: <path to file in S3 after [s3.bucket]/[s3.basepath]>
     filename: <local path + name of the file, will be downloaded into [downloadpath]/[localname]>
     permission: <[optional] file permissions, default 0644>
 - commands:
   <YAML array containing shell commands (currently /bin/bash) to be run before 'docker build'>

 You can use ${<KEY>} as a placeholder in harbor.yml to be replaced by the value passed in a -e flag

Options:
  -e []: List of KEY=VALUE to be replaced in harbor.yml
  -v: Display version information
  -no-download: do not download files
  -no-command: do not run commands
  -no-docker: do not run 'docker build', 'docker tag' and 'docker push' after file downloads and command runs
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
