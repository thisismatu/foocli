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

func showInfo() string {
	icon := `💬`
	name := "Foo CLI"
	currVersion := "v0.0.1"
	nextVersion := "v0.0.2"
	nameAndVersion := fmt.Sprintf("\n%s %s %s", icon, name, currVersion)
	if currVersion != nextVersion {
		update := fmt.Sprintf("> Update available: run %s to install %s %s\n", color.CyanString("brew uprade foo"), name, nextVersion)
		return update + nameAndVersion
	}
	return nameAndVersion
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
	rootCmd.AddCommand(modelsCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(transcribeCmd)
}
