package cmd

import (
	"github.com/spf13/cobra"
)

var modelsCmd = &cobra.Command{
	Use:     "models",
	Aliases: []string{"model"},
	Short:   "View and manage models",
	Long:    "List, view and manage models",
	Args:    cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			modelsListCmd.Run(cmd, []string{})
		}
		if len(args) > 0 {
			modelsInfoCmd.Run(cmd, args)
		}
	},
}

func init() {
}
