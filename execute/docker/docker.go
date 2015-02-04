package docker

import (
	"fmt"
	"github.com/victorcampos/harbor/execute"
	"os"
	"time"
)

func Build(imageTag string, tags []string) error {
	cwd, _ := os.Getwd()

	// FIXME: Use Docker Remote API to perform docker commands
	// Current usage is too OS-specific and depends on sudo power when executing harbor
	if err := execute.CommandWithArgs("docker", "build", "-t", imageTag, cwd); err != nil {
		return err
	}

	validatedTags := buildTags(tags)
	for _, tag := range validatedTags {
		if err := createAndPushTag(imageTag, tag); err != nil {
			return err
		}
	}

	return nil
}

func buildTags(tags []string) []string {
	// always put latest
	validatedTags := []string{"latest"}

	if tags != nil {
	for _, value := range tags {
			if value == "latest" { continue }
			validatedTags = append(validatedTags, value)
		}
	}

	// Default version if not set
	validatedTagsLen := len(validatedTags)

	if validatedTagsLen == 1 {
		// default version format: the current date time in format YYYYMMDDTHHMM
		version := time.Now().Format("20060102T1504")
		validatedTags = append(validatedTags, version)
	}

	return validatedTags
}

func createAndPushTag(imageTag string, version string) error {
	versionedImageTag := fmt.Sprintf("%s:%s", imageTag, version)

	if err := execute.CommandWithArgs("docker", "tag", "-f", imageTag, versionedImageTag); err != nil {
		return err
	}

	if err := execute.CommandWithArgs("docker", "push", versionedImageTag); err != nil {
		return err
	}

	return nil
}
