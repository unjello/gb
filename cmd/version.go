package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/unjello/gb/core"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gB",
	Long:  `All software has versions. This is gB's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(core.Desc, "v"+core.Version)
	},
}
