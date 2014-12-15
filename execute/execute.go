package execute

import (
	"fmt"
	"github.com/victorcampos/harbor/config"
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

			command := exec.Command(commandHead, commandArgs...)
			commandOutput, err := command.CombinedOutput()

			if err != nil {
				return err
			}

			fmt.Printf("--- Output of: %s %s\r\n%s", commandHead, commandArgs, string(commandOutput))
		}
	}

	return nil
}

func splitCommand(command string) (commandHead string, commandArgs []string) {
	commandParts := strings.Fields(command)
	commandHead = commandParts[0]
	commandArgs = commandParts[1:len(commandParts)]

	return
}
