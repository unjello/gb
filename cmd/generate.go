package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/evilsocket/islazy/log"
	"github.com/evilsocket/islazy/tui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

func checkIfSourceFolderExists(project_root string) error {
	path := filepath.Join(project_root, "src")
	fi, err := os.Stat(path)
	if err != nil {
		log.Warning("Source folder " + tui.Green("src") + " not found in project root. You should create one.")
		return err
	}
	if mode := fi.Mode(); mode.IsDir() != true {
		log.Warning(tui.Green("src") + " found in project root, but it is not a folder.")
		return fmt.Errorf("Source folder is not a directory")
	}

	return nil
}

func checkIfBuildFolderIsIgnored(project_root string) error {
	path := filepath.Join(project_root, ".gitignore")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Warning("No " + tui.Green(".gitignore") + " in project root. You should create one.")
		return err
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Could not read .gitignore contents")
	}
	r, _ := regexp.Compile("(?m)^build/?$")
	match := r.Match(b)
	if match == false {
		log.Warning("Your build folder " + tui.Green("build/") + " should be ignored by git. Add it to " + tui.Green(".gitignore"))
	}
	return nil
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate build files",
	Long:  `Generate build files`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal("Could not get current working directory")
		}
		log.Info("Generating build for project dir: " + tui.Dim(cwd))
		if checkIfBuildFolderIsIgnored(cwd) != nil {
			return
		}
		if checkIfSourceFolderExists(cwd) != nil {
			return
		}
	},
}
