package docker

import (
	"testing"
	"time"
)

func TestBuildTagsShouldAddLatestWhenEmpty(t *testing.T) {
	emptyTags := make([]string, 0)

	returnedTags := buildTags(emptyTags)
	if returnedTags[0] != "latest" {
		t.Error("Expected 'latest' in empty tags slice")
	}
}

func TestBuildTagsShouldNotAddLatestIfPresent(t *testing.T) {
	tagsWithLatest := make([]string, 1)
	tagsWithLatest[0] = "latest"

	returnedTags := buildTags(tagsWithLatest)
	latestTagCount := 0

	for _, tag := range returnedTags {
		if tag == "latest" {
			latestTagCount++
		}
	}

	if latestTagCount != 1 {
		t.Error("Expected 'latest' once in tags slice, found", latestTagCount)
	}
}

func TestCreateTimeBasedVersion(t *testing.T) {
	myTime := time.Date(2006, time.January, 2, 15, 4, 0, 0, time.UTC)

	version := createTimeBasedVersion(myTime)

	if version != "20060102T1504" {
		t.Error("Expected version 20060102T1504, got", version)
	}
}
