package main

import (
	"fmt"
	"github.com/victorcampos/harbor/config"
	"github.com/victorcampos/harbor/download"
	"github.com/victorcampos/harbor/execute"
	"os"
)

func main() {
	harborConfig, err := config.Load()
	checkError(err)

	err = download.FromS3(harborConfig)
	checkError(err)

	err = execute.Commands(harborConfig)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
