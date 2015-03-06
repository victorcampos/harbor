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
	buildImageTag := imageTag

	if config.Options.NoLatestTag && len(supplementaryTags) > 0 {
		// Always create the build image tag from the first supplied tag if -no-latest-tag is set
		buildImageTag = versionedImageTag(imageTag, supplementaryTags[0])
	} else if config.Options.NoLatestTag && len(supplementaryTags) == 0 {
		// Else create a time-based version and avoid creating 'latest' tag
		version := createTimeBasedVersion(time.Now())
		buildImageTag = versionedImageTag(imageTag, version)
		supplementaryTags = append(supplementaryTags, version)
	}

	// FIXME: Use Docker Remote API to perform docker commands
	// Current usage is too OS-specific and depends on sudo power when executing harbor
	if err := runDockerCommand("build", "-t", buildImageTag, cwd); err != nil {
		return err
	}

	finalTags := buildTags(supplementaryTags)
	for _, tag := range finalTags {
		if err := createTag(buildImageTag, imageTag, tag); err != nil {
			return err
		}

		if !config.Options.NoDockerPush {
			if err := pushTag(imageTag, tag); err != nil {
				return err
			}
		}
	}

	return nil
}

func buildTags(tags []string) []string {
	finalTags := []string{}

	if tags != nil {
		for _, tag := range tags {
			if tag != "latest" {
				finalTags = append(finalTags, tag)
			}
		}
	}

	// Adds default version if no additional tag given and -no-latest-tag is unset
	if len(tags) == 0 && !config.Options.NoLatestTag {
		// default version format: the current date time in format YYYYMMDDTHHMM
		version := createTimeBasedVersion(time.Now())
		finalTags = append(finalTags, version)
	}

	return finalTags
}

func createTimeBasedVersion(t time.Time) string {
	return t.Format("20060102T1504")
}

func createTag(fromTag string, imageTag string, version string) error {
	versionedImageTag := versionedImageTag(imageTag, version)

	if err := runDockerCommand("tag", "-f", fromTag, versionedImageTag); err != nil && fromTag != versionedImageTag {
		return err
	}

	return nil
}

func pushTag(imageTag string, version string) error {
	versionedImageTag := versionedImageTag(imageTag, version)

	if err := runDockerCommand("push", versionedImageTag); err != nil {
		return err
	}

	return nil
}

func versionedImageTag(imageTag string, version string) string {
	return fmt.Sprintf("%s:%s", imageTag, version)
}

func runDockerCommand(parameters ...string) error {
	if len(config.Options.DockerOpts) > 0 {
		parameters = append([]string{config.Options.DockerOpts}, parameters...)
	}

	return execute.CommandWithArgs("docker", parameters...)
}
