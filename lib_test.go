package examplegolibrary

import (
	"testing"
)

var expectedVersion = "v1.0.0"

func TestVersion(t *testing.T) {
	if expectedVersion != Version {
		t.Errorf("Unexpected version number from Version: expected %s, got %s", expectedVersion, Version)
	}
}

func TestGetVersion(t *testing.T) {
	if expectedVersion != GetVersion() {
		t.Errorf("Unexpected version number from GetVersion: expected %s, got %s", expectedVersion, GetVersion())
	}
}
