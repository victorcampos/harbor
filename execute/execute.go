package execute

import (
	"fmt"
	"github.com/victorcampos/harbor/config"
	"os"
	"os/exec"
	"strings"
)

func Commands(harborConfig config.HarborConfig) error {
	commandListLength := len(harborConfig.Commands)

	if commandListLength > 0 {
		fmt.Printf("--- Commands to be executed: %d\r\n", commandListLength)
		for key, value := range harborConfig.Commands {
			fmt.Printf("--- Executing command %d of %d...\r\n", key+1, commandListLength)

			commandHead, commandArgs := splitCommand(value)
			fmt.Printf("--- Executing %s with \"%s\"\r\n", commandHead, commandArgs)

			if err := CommandWithArgs(commandHead, commandArgs...); err != nil {
				return err
			}
		}
	}

	return nil
}

func splitCommand(command string) (commandHead string, commandArgs []string) {
	commandParts := strings.Fields(command)
	commandHead = commandParts[0]
	commandArgs = commandParts[1:]

	return
}

func CommandWithArgs(commandHead string, commandArgs ...string) error {
	args := strings.Join(commandArgs, " ")
	concatenatedCommand := strings.Join([]string{commandHead, args}, " ")

	command := exec.Command("/bin/bash", "-c", concatenatedCommand)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()

	return err
}
