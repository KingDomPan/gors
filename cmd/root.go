package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "gors",
	Short: "gors is a terminal screen tools for record and play",
	Long:  "just use `record` and `play` cmd",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(cmd.UsageString())
		}
	},
}
