package cmd

import (
	"github.com/spf13/cobra"
)

var modelsCmd = &cobra.Command{
	Use:     "models",
	Aliases: []string{"model"},
	Short:   "View and manage models",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
