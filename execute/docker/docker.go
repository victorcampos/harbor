package docker

import (
	"fmt"
	"github.com/victorcampos/harbor/config"
	"github.com/victorcampos/harbor/execute"
	"os"
	"time"
)

func Build(imageName string, tags []string) error {
	cwd, _ := os.Getwd()

	if len(tags) == 0 {
		timeBasedTag := createTimeBasedVersion(time.Now())
		tags = append([]string{timeBasedTag}, tags...)
	}

	if !config.Options.NoLatestTag {
		tags = append([]string{"latest"}, tags...)
	}

	imageWithTagsList := createImageWithTagsList(imageName, tags)

	if err := runDockerCommand("build", "-t", imageWithTagsList[0], cwd); err != nil {
		return err
	}

	for i, taggedImage := range imageWithTagsList {
		if i != 0 {
			if err := createTag(imageWithTagsList[0], taggedImage); err != nil {
				return err
			}
		}

		if !config.Options.NoDockerPush {
			if err := pushTag(taggedImage); err != nil {
				return err
			}
		}
	}

	return nil
}

func createImageWithTagsList(imageName string, tags []string) []string {
	var imageWithTagsList []string

	for _, tag := range tags {
		imageWithTagsList = append(imageWithTagsList, concatenateImageWithTag(imageName, tag))
	}

	return imageWithTagsList
}

func concatenateImageWithTag(imageName string, tag string) string {
	return imageName + ":" + tag
}

func createTimeBasedVersion(t time.Time) string {
	return t.Format("20060102T1504")
}

func createTag(fromTag string, toTag string) error {
	if err := runDockerCommand("tag", "-f", fromTag, toTag); err != nil {
		return err
	}

	return nil
}

func pushTag(taggedImage string) error {
	if err := runDockerCommand("push", taggedImage); err != nil {
		return err
	}

	return nil
}

func runDockerCommand(parameters ...string) error {
	if len(config.Options.DockerOpts) > 0 {
		parameters = append([]string{config.Options.DockerOpts}, parameters...)
	}

	if !config.Options.Debug {
		return execute.CommandWithArgs("docker", parameters...)
	} else {
		fmt.Printf("docker %s\n", parameters)
		return nil
	}
}
