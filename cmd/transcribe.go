package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var transcribeCmd = &cobra.Command{
	Use:   "transcribe",
	Short: "Transcribe audio files",
	Example: `  foo transcribe file.wav
  foo transcribe files.jsonl <model_id>`,
	Args: cobra.RangeArgs(0, 2),
	Run: func(cmd *cobra.Command, args []string) {
		mid, err := cmd.Flags().GetString("model")
		if err != nil {
			logError(err)
		}
		if mid == "" && len(args) > 0 {
			mid = args[0]
		}
		input, err := cmd.Flags().GetString("model")
		if err != nil {
			logError(err)
		}
		if input == "" && len(args) > 1 {
			input = args[1]
		}

		fmt.Printf("%s %s\n", mid, input)
	},
}

func init() {
	transcribeCmd.Flags().StringP("model", "m", "", "Specify which model to use")
	transcribeCmd.Flags().StringP("input", "i", "", "File(s) to transcribe")
}
