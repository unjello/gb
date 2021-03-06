package cmd

import (
	"github.com/spf13/cobra"
	"github.com/unjello/gb/core"
)

func init() {
	rootCmd.AddCommand(makeCmd)
}

var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "Generate & Build",
	Long:  "Generate & Build",
	Run: func(cmd *cobra.Command, args []string) {
		core.GenerateBuildScripts()
		core.BuildProject()
		core.RunTests()
	},
}
