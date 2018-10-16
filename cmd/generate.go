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

func _check_if_src_folder_exists(project_root string) error {
	src_path := filepath.Join(project_root, "src")
	fi, err := os.Stat(src_path)
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

func _check_if_build_folder_is_ignored(project_root string) error {
	gitignore_path := filepath.Join(project_root, ".gitignore")
	if _, err := os.Stat(gitignore_path); os.IsNotExist(err) {
		log.Warning("No " + tui.Green(".gitignore") + " in project root. You should create one.")
		return err
	}

	b, err := ioutil.ReadFile(gitignore_path)
	if err != nil {
		log.Fatal("Could not read .gitignore contents")
	}
	r, _ := regexp.Compile("^build/?$")
	match := r.Find(b)
	if match == nil {
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
		if _check_if_build_folder_is_ignored(cwd) != nil {
			return
		}
		if _check_if_src_folder_exists(cwd) != nil {
			return
		}
	},
}
