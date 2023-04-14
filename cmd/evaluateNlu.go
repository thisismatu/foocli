package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var evaluateNluCmd = &cobra.Command{
	Use:   "nlu [path] [id]",
	Short: "Evaluate model NLU accuracy",
	Long:  fmt.Sprintf("Evaluates NLU accuracy for the given model.\n\nTo run NLU evaluation, you need a set of ground truth annotations.\nUse the %s command to get started, or read the docs:\nhttps://docs.foo.com/evaluate-nlu", fmtCmd("annotate")),
	Example: `  Basic evaluation
  $ foo evaluate nlu ground-truths.txt my-model-id

  Evaluate using a reference date
  $ foo evaluate nlu ground-truths.txt my-model-id --date 2021-01-20`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	evaluateNluCmd.Flags().StringP("date", "d", "", "Reference date in YYYY-MM-DD format")
}
