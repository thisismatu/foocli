package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var modelCmd = &cobra.Command{
	Use:     "model",
	Aliases: []string{"models"},
	Short:   "View, deploy and manage models",
	Args:    cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	modelCmd.AddCommand(modelListCmd)
	modelCmd.AddCommand(modelInfoCmd)
	modelCmd.AddCommand(modelAddCmd)
	modelCmd.AddCommand(modelRemoveCmd)
	modelCmd.AddCommand(modelDeployCmd)
	modelCmd.AddCommand(modelDownloadCmd)
}
