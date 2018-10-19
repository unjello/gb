package core

import (
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"

	"github.com/evilsocket/islazy/log"
)

const (
	errorNoCommand             = "No command to execute"
	debugExecutingCommand      = "Executing command: %s"
	fatalFailedToGetStderrPipe = "Failed to get stderr pipe"
	fatalFailedToGetStdoutPipe = "Failed to get stdout pipe"
	fatalFailedToStartCommand  = "Failed to start command"
	fatalCommandFailedToRun    = "Command failed to run"
)

func PrintCommand(command []string, isDebug bool) error {
	stdout, stderr, err := RunCommand(command)

	if isDebug {
		if len(stderr) > 0 {
			fmt.Println(stderr)
		}
		if len(stdout) > 0 {
			fmt.Println(stdout)
		}
	}

	if err != nil {
		if isDebug == false && len(stderr) > 0 {
			fmt.Println(stderr)
		}
		return err
	}

	return nil
}

func RunCommand(command []string) (string, string, error) {
	var (
		stdout io.ReadCloser
		stderr io.ReadCloser
		err    error
	)

	if len(command) < 1 {
		return "", "", fmt.Errorf(errorNoCommand)
	}
	log.Debug(debugExecutingCommand, command)

	c := exec.Command(command[0], command[1:]...)

	if stderr, err = c.StderrPipe(); err != nil {
		log.Fatal(fatalFailedToGetStderrPipe)
		return "", "", err
	}
	if stdout, err = c.StdoutPipe(); err != nil {
		log.Fatal(fatalFailedToGetStdoutPipe)
		return "", "", err
	}

	if err = c.Start(); err != nil {
		log.Fatal(fatalFailedToStartCommand)
		return "", "", err
	}

	stderrOut, _ := ioutil.ReadAll(stderr)
	stdoutOut, _ := ioutil.ReadAll(stdout)

	if err = c.Wait(); err != nil {
		log.Fatal(fatalCommandFailedToRun)
		return string(stdoutOut), string(stderrOut), err
	}

	return string(stdoutOut), string(stderrOut), nil
}
