package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var cfgFile = ".cfg/foo"
var dbProjects = ".cfg/projects.jsonl"
var dbApplications = ".cfg/applications.jsonl"

var rootCmd = &cobra.Command{
	Use:   "foo",
	Short: "Foo CLI",
	Long:  `Foo is a CLI tool for using the Foo platform`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
