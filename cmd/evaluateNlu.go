package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var evaluateNluCmd = &cobra.Command{
	Use:   "nlu [path] [id]",
	Short: "Evaluate model NLU accuracy",
	Long: fmt.Sprintf(`Evaluates NLU accuracy for the given model. To run NLU evaluation, you need a
set of ground truth annotations. Use the %s command to get started.
Check out the docs: https://docs.foo.com/evaluate-nlu`, fmtCmd("annotate")),
	Example: fmtExample("Basic evaluation", "foo evaluate nlu ground-truths.txt my-model-id", false) +
		fmtExample("Evaluate using a reference date", "foo evaluate nlu ground-truths.txt my-model-id --date 2021-01-20", true),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	evaluateNluCmd.Flags().StringP("date", "d", "", "Reference date in YYYY-MM-DD format")
}
