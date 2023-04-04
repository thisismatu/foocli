package cmd

import (
	"github.com/spf13/cobra"
)

var transcribeCmd = &cobra.Command{
	Use:   "transcribe",
	Short: "Transcribe audio files",
	Example: `  foo transcribe file.wav
  foo transcribe files.jsonl -m <model_id>`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	transcribeCmd.Flags().StringP("model", "m", "large-highaccuracy", "Specify which model to use")
}
