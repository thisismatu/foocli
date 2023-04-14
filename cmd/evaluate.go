package cmd

import (
	"github.com/spf13/cobra"
)

var evaluateCmd = &cobra.Command{
	Use:   "evaluate",
	Short: "Evaluate model accuracy",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	evaluateCmd.AddCommand(evaluateAsrCmd)
	evaluateCmd.AddCommand(evaluateNluCmd)
}
