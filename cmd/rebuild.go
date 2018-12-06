package cmd

import (
	"github.com/spf13/cobra"
	"github.com/unjello/gb/core"
)

func init() {
	rootCmd.AddCommand(rebuildCmd)
}

var rebuildCmd = &cobra.Command{
	Use:   "rebuild",
	Short: "ReBuild",
	Long:  "ReBuild",
	Run: func(cmd *cobra.Command, args []string) {
		core.RemoveBuildFiles()
		core.GenerateBuildScripts()
		core.BuildProject()
		core.RunTests()
	},
}
