package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bmatcuk/doublestar"
	"github.com/evilsocket/islazy/log"
	"github.com/evilsocket/islazy/tui"
)

func RunTests() error {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Could not get current working directory")
	}
	buildRoot := filepath.Join(cwd, "build")
	log.Info("Running tests for project: " + tui.Dim(cwd))
	log.Info("Using build folder: " + tui.Dim(buildRoot))
	binRoot := filepath.Join(buildRoot, "bin")
	testsBinRoot := filepath.Join(binRoot, "tests")
	testsGlob := filepath.Join(testsBinRoot, "**/*")
	files, err := doublestar.Glob(testsGlob)
	if err != nil {
		log.Fatal("Failed to find tests using pattern: " + tui.Red(testsGlob))
		log.Fatal(err.Error())
		return err
	}

	runner := OsCommandRunner{}
	anyFailed := false
	for _, file := range files {
		err := runner.RunWithOutput([]string{file})
		if err != nil {
			anyFailed = true
		}
	}

	if anyFailed {
		return fmt.Errorf("Some tests failed")
	}

	return nil
}
