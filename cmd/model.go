package cmd

import (
	"github.com/spf13/cobra"
)

var modelCmd = &cobra.Command{
	Use:     "model",
	Aliases: []string{"models"},
	Short:   "View and manage models",
	Long:    "List, view and manage models",
	Args:    cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			modelListCmd.Run(cmd, []string{})
		}
		if len(args) > 0 {
			modelInfoCmd.Run(cmd, args)
		}
	},
}

func init() {
}
