package core

import (
	"os/exec"
	"testing"
)

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	return exec.Command("go")
}

func TestRunCommandWithNoCommandFails(t *testing.T) {
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()

	if err := runCommand(nil, false); err == nil {
		t.Errorf("Expected error, got %q", err)
	}
}
