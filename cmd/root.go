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
	logoAndVersion := color.CyanString("\n%s %s", logo, currVersion)
	docs := color.New(color.Faint).Sprint("\nhttps://docs.foo.com/cli")
	nextVersion := "0.0.2"
	if currVersion != nextVersion {
		update := fmt.Sprintf("%s run %s to install version %s\n", color.MagentaString("Update available:"), color.CyanString("`brew uprade foo`"), nextVersion)
		return update + logoAndVersion + docs
	}
	return logoAndVersion + docs
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
}
