package core

import (
	"os"
	"path/filepath"

	"github.com/evilsocket/islazy/log"
	"github.com/evilsocket/islazy/tui"
)

func BuildProject() {
	if err := VerifyNinjaExists(); err != nil {
		log.Fatal("Could not find Ninja. Please install it from https://ninja-build.org/")
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Could not get current working directory")
	}
	buildRoot := filepath.Join(cwd, "build")
	log.Info("Building project: " + tui.Dim(cwd))
	log.Info("Using build folder: " + tui.Dim(buildRoot))
	runner := OsCommandRunner{}
	runner.RunWithOutput([]string{"ninja", "-C", "build"})
}
