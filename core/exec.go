package core

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/evilsocket/islazy/log"
)

const (
	errorNoCommand          = "No command to execute"
	debugExecutingCommand   = "Executing command: %s"
	fatalCommandFailedToRun = "Command failed to run"
)

/*
   this little trick for testing exec.Command taken from
   https://npf.io/2015/06/testing-exec-command/
*/
var execCommand = exec.Command

func runCommand(command []string, showOutput bool) error {
	if len(command) < 1 {
		return fmt.Errorf(errorNoCommand)
	}
	log.Debug(debugExecutingCommand, command)

	c := execCommand(command[0], command[1:]...)

	if showOutput {
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
	}

	if err := c.Run(); err != nil {
		log.Error(fatalCommandFailedToRun)
		return err
	}

	return nil
}

type CommandRunner interface {
	Run(command []string) error
	RunWithOutput(command []string) error
}
type OsCommandRunner struct{}

func (OsCommandRunner) Run(command []string) error           { return runCommand(command, false) }
func (OsCommandRunner) RunWithOutput(command []string) error { return runCommand(command, true) }

func NewOsCommandRunner() CommandRunner { return OsCommandRunner{} }
