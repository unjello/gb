package core

import (
	"testing"
)

func TestExecuteTemplateDoesNotEscapeStrings(t *testing.T) {
	buildInfo := NewBuildInfo()
	buildInfo.Cxx = "g++"
	ninjaTemplate := "{{ .Cxx }}"
	ninjaFile, err := ExecuteTemplate("ninjaBuildTemplate", ninjaTemplate, buildInfo)

	if err != nil {
		t.Errorf("Expected no error, got: %q", err)
	}

	if ninjaFile != "g++" {
		t.Errorf("Expected name to be 'g++', got: '%s'", ninjaFile)
	}
}
