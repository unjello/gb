package cmd

import (
	"github.com/spf13/cobra"
	"github.com/unjello/gb/core"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate build files",
	Long:  `Generate build files`,
	Run: func(cmd *cobra.Command, args []string) {
		core.GenerateBuildScripts()
	},
}
