package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/victorcampos/harbor/commandline"
	"github.com/victorcampos/harbor/config"
	"github.com/victorcampos/harbor/download"
	"github.com/victorcampos/harbor/execute"
	"github.com/victorcampos/harbor/execute/docker"
	"os"
)

const VERSION = "0.2.0"

func main() {
	usage := `Harbor, a Docker wrapper

Harbor looks up a file named harbor.yml in the same directory where run from, harbor.yml structure is:
 imagetag: <tag to be used on 'docker build'>
 tags:
   - <YAML array of custom tags to create and push into registry>
 downloadpath: <local root path to download files into>
 s3:
   bucket: <base bucket to download files from>
   basepath: <inside the bucket the root path for files to be downloaded>
   region: <[optional] region of the bucket, default us-east-1>
 files:
   - s3path: <path to file in S3 after [s3.bucket]/[s3.basepath]>
     filename: <local path + name of the file, will be downloaded into [downloadpath]/[localname]>
     permission: <[optional] file permissions, default 0644>
 commands:
   - <YAML array containing shell commands (currently /bin/bash) to be run before 'docker build'>

 You can use ${<KEY>} as a placeholder in harbor.yml to be replaced by the value passed in a -e flag

Usage:
  harbor -h | --help
  harbor --version
  harbor --list-variables
  harbor [-e KEY=VALUE]... [options]
  harbor [options]

Options:
  -h, --help                    Show this screen.
  -v, --version                 Show version.
  --list-variables         Parses harbor.yml and prints out every ${KEY} found.
  -e KEY=VALUE                  Replaces every ${KEY} in harbor.yml with VALUE
  --debug                       Dry-run and print command executions.
  --no-download                 Prevents downloading files from S3.
  --no-commands                 Prevents commands in harbor.yml from being run.
  --no-docker                   Do not run 'docker build', 'docker tag' and 'docker push' commands.
  --no-docker-push              Do not run 'docker push' after building and tagging images.
  --docker-opts=<DOCKER_OPTS>   Will be appended to the base docker command (ex.: 'docker <docker-opts> command').
  --no-latest-tag               Do not build image tagging as 'latest',
                                will use the first tag in 'tags' list from harbor.yml or
                                generate a timestamped tag if no 'tags' list exists.`

	arguments, _ := docopt.Parse(usage, nil, true, "Harbor "+VERSION, true)

	configVars := arguments["-e"].([]string)
	debugFlag := arguments["--debug"].(bool)
	listVariablesFlag := arguments["--list-variables"].(bool)
	noDownloadFlag := arguments["--no-download"].(bool)
	noCommandFlag := arguments["--no-commands"].(bool)
	noDockerFlag := arguments["--no-docker"].(bool)
	noDockerPushFlag := arguments["--no-docker-push"].(bool)
	dockerOpts, _ := arguments["--docker-opts"].(string)
	noLatestTagFlag := arguments["--no-latest-tag"].(bool)

	if listVariablesFlag {
		listVariables()
	}

	cliConfigVars, err := commandline.NewConfigVarsMap(configVars)
	if err != nil {
		checkError(err)
	}

	harborConfig, err := config.Load(cliConfigVars)
	checkError(err)

	config.Options.Debug = debugFlag
	config.Options.DockerOpts = dockerOpts
	config.Options.NoDockerPush = noDockerPushFlag
	config.Options.NoLatestTag = noLatestTagFlag

	if !noDownloadFlag {
		err = download.FromS3(harborConfig)
		checkError(err)
	}

	if !noCommandFlag {
		err = execute.Commands(harborConfig)
		checkError(err)
	}

	if !noDockerFlag {
		err = docker.Build(harborConfig.ImageTag, harborConfig.Tags)
		checkError(err)
	}
}

func listVariables() {
	harborConfigFile, err := config.LoadFile()
	if err != nil {
		checkError(err)
	}

	variablesFound := config.ReadEnv(harborConfigFile)

	fmt.Printf("--- Found %d variables in harbor.yml\n", len(variablesFound))

	for _, variable := range variablesFound {
		fmt.Printf("---   Found: %s\n", variable)
	}

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
