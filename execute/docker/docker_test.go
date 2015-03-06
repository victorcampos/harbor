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
