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

func PrintCommand(command []string, isDebug bool) error {
	return RunCommand(command)
}

func RunCommand(command []string) error {
	if len(command) < 1 {
		return fmt.Errorf(errorNoCommand)
	}
	log.Debug(debugExecutingCommand, command)

	c := exec.Command(command[0], command[1:]...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		log.Fatal(fatalCommandFailedToRun)
		return err
	}

	return nil
}
