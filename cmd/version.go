package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gors",
	Long:  `All software has versions. This is Gors's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gors Version v0.1 -- HEAD")
	},
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}
