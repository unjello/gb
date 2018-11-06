package cmd

import (
	"github.com/spf13/cobra"
	"github.com/unjello/gb/core"
)

func init() {
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run tests",
	Long:  "Run tests",
	Run: func(cmd *cobra.Command, args []string) {
		core.RunTests()
	},
}
