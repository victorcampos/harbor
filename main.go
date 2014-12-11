package main

import (
	"fmt"
	"github.com/victorcampos/harbor/config"
	"github.com/victorcampos/harbor/downloader"
	"os"
	"os/exec"
	"strings"
)

func main() {
	harborConfig, err := config.Load()
	checkError(err)

	err = downloader.DownloadFromS3(harborConfig)
	checkError(err)

	commandListLength := len(harborConfig.Commands)

	if commandListLength > 0 {
		fmt.Printf("Commands to be executed: %d\r\n\r\n", commandListLength)
		for key, value := range harborConfig.Commands {
			fmt.Printf("Executing command number %d of %d...\r\n", key+1, commandListLength)

			commandParts := strings.Fields(value)
			commandHead := commandParts[0]
			commandArgs := commandParts[1:len(commandParts)]

			command := exec.Command(commandHead, commandArgs...)
			commandOutput, _ := command.CombinedOutput()

			fmt.Printf("Executing: %s\r\n", commandHead)
			fmt.Printf("Output of: %s\r\n%s", commandHead, string(commandOutput))
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
