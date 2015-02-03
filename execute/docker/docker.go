package docker

import (
	"fmt"
	"github.com/victorcampos/harbor/execute"
	"os"
	"time"
)

func Build(imageTag string) error {
	cwd, _ := os.Getwd()
	version := time.Now().Format("20060102T1504")

	// Appends the current date time in format YYYYMMDD HHMM as image version
	versionedImageTag := fmt.Sprintf("%s:%s", imageTag, version)

	// FIXME: Use Docker Remote API to perform docker commands
	// Current usage is too OS-specific and depends on sudo power when executing harbor
	if err := execute.CommandWithArgs("docker", "build", "-t", versionedImageTag, cwd); err != nil {
		return err
	}

	if err := execute.CommandWithArgs("docker", "tag", versionedImageTag, fmt.Sprintf("%s:latest", imageTag)); err != nil {
		return err
	}

	if err := push(versionedImageTag); err != nil {
		return err
	}

	if err := push(fmt.Sprintf("%s:latest", imageTag)); err != nil {
		return err
	}

	return nil
}

func push(imageTag string) error {
	if err := execute.CommandWithArgs("docker", "push", imageTag); err != nil {
		return err
	}

	return nil
}
