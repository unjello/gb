package core

import (
	"testing"
)

func TestGetPackageNameCorrectlyParsesUri(t *testing.T) {
	pkg, err := ParsePackageString("doctest/2.0.0@unjello/testing")
	if err != nil {
		t.Errorf("Expected no error, got: %q", err)
	}

	if pkg.Name != "doctest" {
		t.Errorf("Expected name to be doctest, got: %s", pkg.Name)
	}
	if pkg.Version != "2.0.0" {
		t.Errorf("Expected version to be 2.0.0, got: %s", pkg.Version)
	}
	if pkg.Channel != "unjello/testing" {
		t.Errorf("Expected channel to be unjello/testing, got: %s", pkg.Channel)
	}
}
