package cmd

import (
	"github.com/spf13/cobra"
	"github.com/unjello/gb/core"
)

func init() {
	rootCmd.AddCommand(cleanCmd)
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean",
	Long:  "Clean",
	Run: func(cmd *cobra.Command, args []string) {
		core.RemoveBuildFiles()
	},
}
