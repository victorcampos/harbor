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

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
