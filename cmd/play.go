package cmd

import (
	"fmt"

	"github.com/KingDomPan/gors/play"

	"github.com/spf13/cobra"
)

var pfilename string

var PlayCmd = &cobra.Command{
	Use:   "play",
	Short: "play you terminal by a filename",
	Long:  "play you terminal by a filename",
	Run: func(cmd *cobra.Command, args []string) {
		if pfilename == "" {
			fmt.Println(cmd.UsageString())
			return
		}
		player := &play.Player{
			Filename: pfilename,
		}
		player.Execute()
	},
}

func init() {
	RootCmd.AddCommand(PlayCmd)
	PlayCmd.Flags().StringVarP(&pfilename, "filename", "f", "", "the file which to play your terminal data")
}
