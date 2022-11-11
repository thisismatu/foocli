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

var rootCmd = &cobra.Command{
	Use:     "foo",
	Short:   "Foo CLI",
	Long:    "Foo is a CLI tool for using the Foo platform",
	Example: fmt.Sprintf("  Deploy changes\n  %s\n\n  Transcribe audio file\n  %s\n\n  Switch project\n  %s\n", color.CyanString("foo deploy <app_id> /path/to/config"), color.CyanString("foo transcribe <app_id> file.wav"), color.CyanString("foo projects switch")),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
