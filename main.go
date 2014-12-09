package main

import (
	"fmt"
	loader "github.com/victorcampos/harbor/loader"
	"os/exec"
	"strings"
)

func main() {
	harborConfig, err := loader.LoadConfig()

	if err != nil {
		fmt.Println(err)
	}

	/*for key, value := range harborConfig.Files {
	}*/

	commandListCount := len(harborConfig.Commands)

	fmt.Printf("Commands to be executed: %d\r\n\r\n", commandListCount)
	for key, value := range harborConfig.Commands {
		fmt.Printf("Executing command number %d of %d...\r\n", key+1, commandListCount)

		commandParts := strings.Fields(value)
		commandHead := commandParts[0]
		commandParts = commandParts[1:len(commandParts)]

		command := exec.Command(commandHead, commandParts...)
		commandOutput, _ := command.CombinedOutput()

		fmt.Printf("Executing: %s\r\n", commandHead)
		fmt.Printf("Output of: %s\r\n%s", commandHead, string(commandOutput))
	}
}
