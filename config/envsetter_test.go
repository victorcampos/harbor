package config

import (
	"testing"
)

func TestReadEnv(t *testing.T) {
	configString := []byte("key: ${SOMEKEY}\nanother: ${ANOTHERKEY}")

	varsFound := ReadEnv(configString)

	if len(varsFound) != 2 {
		t.Fatal("Expected 2 variables found, got", len(varsFound))
	}

	if varsFound[0] != "${SOMEKEY}" {
		t.Fatal("Expected ${SOMEKEY}, got", varsFound[0])
	}

	if varsFound[1] != "${ANOTHERKEY}" {
		t.Fatal("Expected ${ANOTHERKEY}, got", varsFound[1])
	}
}
