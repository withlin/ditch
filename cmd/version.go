package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/withlin/ditch/ditch"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  `Show the version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Ditch %s\n", ditch.Version)
	},
}

func init() {
	ditchCmd.AddCommand(versionCmd)
}