package docker

import (
	"fmt"
	"github.com/victorcampos/harbor/config"
	"github.com/victorcampos/harbor/execute"
	"os"
	"time"
)

func Build(imageTag string, supplementaryTags []string) error {
	cwd, _ := os.Getwd()

	// FIXME: Use Docker Remote API to perform docker commands
	// Current usage is too OS-specific and depends on sudo power when executing harbor
	if err := runDockerCommand("build", "-t", imageTag, cwd); err != nil {
		return err
	}

	finalTags := buildTags(supplementaryTags)
	for _, tag := range finalTags {
		if err := createAndPushTag(imageTag, tag); err != nil {
			return err
		}
	}

	return nil
}

func buildTags(tags []string) []string {
	// always put latest
	finalTags := []string{"latest"}

	if tags != nil {
		for _, tag := range tags {
			if tag != "latest" {
				finalTags = append(finalTags, tag)
			}
		}
	}

	// Adds default version if no additional tag given
	if len(tags) == 0 {
		// default version format: the current date time in format YYYYMMDDTHHMM
		version := createTimeBasedVersion(time.Now())
		finalTags = append(finalTags, version)
	}

	return finalTags
}

func createTimeBasedVersion(t time.Time) string {
	return t.Format("20060102T1504")
}

func createAndPushTag(imageTag string, version string) error {
	versionedImageTag := fmt.Sprintf("%s:%s", imageTag, version)

	if err := runDockerCommand("tag", "-f", imageTag, versionedImageTag); err != nil {
		return err
	}

	if err := runDockerCommand("push", versionedImageTag); err != nil {
		return err
	}

	return nil
}

func runDockerCommand(parameters ...string) error {
	if len(config.Options.DockerOpts) > 0 {
		parameters = append([]string{config.Options.DockerOpts}, parameters...)
	}

	return execute.CommandWithArgs("docker", parameters...)
}
