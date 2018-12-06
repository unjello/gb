package core

import (
	"os"
	"path/filepath"

	"github.com/evilsocket/islazy/log"
)

func RemoveBuildFiles() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Could not get current working directory")
	}
	buildRoot := filepath.Join(cwd, "build")

	if _, err := os.Stat(buildRoot); os.IsNotExist(err) {
		return
	}

	os.RemoveAll(buildRoot)
}
