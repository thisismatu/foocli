package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var cfgFile = ".cfg/foo"
var dbProjects = ".cfg/projects.jsonl"
var dbModels = ".cfg/models.jsonl"
var currVersion = "0.0.1"

func showInfo() string {
	logo := "ⴲ Foo CLI"
	logoAndVersion := color.CyanString("\n%s %s\n", logo, currVersion)
	docsLink := faint("https://docs.foo.com/cli")
	nextVersion := "0.0.2"
	if currVersion != nextVersion {
		update := fmt.Sprintf("%s run %s to install version %s\n", color.MagentaString("Update available:"), fmtCmd("brew uprade foo"), nextVersion)
		return update + logoAndVersion + docsLink
	}
	return logoAndVersion + docsLink
}

var rootCmd = &cobra.Command{
	Use:     "foo",
	Short:   "Foo CLI",
	Long:    showInfo(),
	Version: currVersion,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(projectCmd)
	rootCmd.AddCommand(modelCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(transcribeCmd)
	rootCmd.AddCommand(evaluateCmd)
	rootCmd.AddCommand(annotateCmd)
	rootCmd.AddCommand(statsCmd)
}
