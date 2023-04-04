package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var cfgFile = ".cfg/foo"
var dbProjects = ".cfg/projects.jsonl"
var dbApplications = ".cfg/applications.jsonl"

func showInfo() string {
	update := fmt.Sprintf("> Update available: run `%s` to install Foo CLI 0.0.2\n\n", color.CyanString("brew uprade foo"))
	version := "💎 Foo CLI 0.0.1"
	return update + version
}

var rootCmd = &cobra.Command{
	Use:   "foo",
	Short: "Foo CLI",
	Long:  showInfo(),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(projectsCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(transcribeCmd)
}
