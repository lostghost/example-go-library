package examplegolibrary

import (
	"testing"
)

var expectedVersion = "v0.2.0"

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
