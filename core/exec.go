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
		log.Fatal(fatalCommandFailedToRun)
		return err
	}

	return nil
}

func RunCommand(command []string) error           { return runCommand(command, false) }
func RunCommandWithOutput(command []string) error { return runCommand(command, true) }
