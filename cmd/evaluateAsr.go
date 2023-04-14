package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var evaluateAsrCmd = &cobra.Command{
	Use:   "asr [path] [id]",
	Short: "Evaluate model ASR accuracy",
	Long:  fmt.Sprintf("Evaluates ASR accuracy for the given model.\n\nTo run ASR evaluation, you need a set of ground truth transcripts.\nUse the %s command to get started, or read the docs:\nhttps://docs.foo.com/evaluate-asr", fmtCmd("transcribe")),
	Example: fmtExample("Basic evaluation", "foo evaluate asr ground-truths.jsonl my-model-id", false) +
		fmtExample("Evaluate using Streaming API", "foo evaluate asr ground-truths.jsonl my-model-id --date 2021-01-20", true),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	evaluateAsrCmd.Flags().BoolP("streaming", "s", false, "Use Streaming API instead of Batch API")
}
