package cmd

import (
	"github.com/KingDomPan/gors/record"

	"github.com/spf13/cobra"
)

var rfilename string

var RecordCmd = &cobra.Command{
	Use:   "record",
	Short: "record you terminal and save to a filename",
	Long:  "record you terminal and save to a filename",
	Run: func(cmd *cobra.Command, args []string) {
		recorder := &record.Recorder{
			Filename: rfilename,
		}
		recorder.Execute()
	},
}

func init() {
	RootCmd.AddCommand(RecordCmd)
	RecordCmd.Flags().StringVarP(&rfilename, "filename", "f", "", "the file which to record your terminal data")
}
