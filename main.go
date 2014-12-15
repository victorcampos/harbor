package main

import (
	"flag"
	"fmt"
	"github.com/victorcampos/harbor/config"
	"github.com/victorcampos/harbor/download"
	"github.com/victorcampos/harbor/execute"
	"os"
)

func main() {
	environment := flag.String("env", "production", "sets the $ENVIRONMENT substitution string")
	flag.Parse()

	fmt.Printf("Using environment: %s\r\n", *environment)

	harborConfig, err := config.Load(*environment)
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
