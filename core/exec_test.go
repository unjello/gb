package core

import (
	"os/exec"
	"testing"

	"github.com/evilsocket/islazy/log"
)

func init() {
	log.Level = log.FATAL + 1
}

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	return exec.Command("go", "version")
}

func fakeExecFailedCommand(command string, args ...string) *exec.Cmd {
	return exec.Command("go_")
}

func TestRunCommandWithNoCommandFails(t *testing.T) {
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()

	if err := runCommand(nil, false); err == nil {
		t.Errorf("Expected error, got %q", err)
	}
}

func TestRunCommandThatExists(t *testing.T) {
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()

	if err := runCommand([]string{"test"}, false); err != nil {
		t.Errorf("Expected no error, got %q", err)
	}
}

func TestRunCommandThatFails(t *testing.T) {
	execCommand = fakeExecFailedCommand
	defer func() { execCommand = exec.Command }()

	if err := runCommand([]string{"test"}, false); err == nil {
		t.Errorf("Expected error, got %q", err)
	}
}
