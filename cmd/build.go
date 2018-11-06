package cmd

import (
	"github.com/spf13/cobra"
	"github.com/unjello/gb/core"
)

func init() {
	rootCmd.AddCommand(buildCmd)
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build",
	Long:  "Build",
	Run: func(cmd *cobra.Command, args []string) {
		core.BuildProject()
		core.RunTests()
	},
}
