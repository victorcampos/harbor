package docker

import (
	"testing"
	"time"
)

func TestCreateTimeBasedVersion(t *testing.T) {
	myTime := time.Date(2006, time.January, 2, 15, 4, 0, 0, time.UTC)

	version := createTimeBasedVersion(myTime)

	if version != "20060102T1504" {
		t.Error("Expected version 20060102T1504, got", version)
	}
}

func TestCreateImageWithTagsList(t *testing.T) {
	imageName := "test"
	tags := []string{"latest", "first-tag", "second-tag"}

	imageWithTagsList := createImageWithTagsList(imageName, tags)

	if len(imageWithTagsList) != 3 {
		t.Fatal("Expected a list of 3 images, got", len(imageWithTagsList))
	}

	if imageWithTagsList[0] != "test:latest" && imageWithTagsList[1] != "test:first-tag" && imageWithTagsList[2] != "test:second-tag" {
		t.Fatal("Expected test:latest, test:first-tag and test:second-tag, got", imageWithTagsList)
	}
}
