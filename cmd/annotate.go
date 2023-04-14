package cmd

import (
	"github.com/spf13/cobra"
)

var annotateCmd = &cobra.Command{
	Use:   "annotate [path]",
	Short: "Annotate example utterances",
	Long:  `Generates SAL annotations for a set of example utterances.`,
	Example: fmtExample("Annotate example utterances using the default model", "foo annotate input.txt", false) +
		fmtExample("Annotate using a specific model", "foo annotate input.txt my-model-id", false) +
		fmtExample("Annotate using a specific model and reference date", "foo annotate input.txt my-model-id --date 2021-01-20", true),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {}
