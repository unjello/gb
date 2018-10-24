package cmd

import (
	"os"
	"path/filepath"

	"github.com/evilsocket/islazy/log"
	"github.com/evilsocket/islazy/tui"
	"github.com/spf13/cobra"
	"github.com/unjello/gb/core"
)

func init() {
	rootCmd.AddCommand(makeCmd)
}

var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "Build",
	Long:  "Build",
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal("Could not get current working directory")
		}
		buildRoot := filepath.Join(cwd, "build")
		log.Info("Building project: " + tui.Dim(cwd))
		log.Info("Using build folder: " + tui.Dim(buildRoot))
		core.PrintCommand([]string{"ninja", "-C", "build"}, true)
	},
}
