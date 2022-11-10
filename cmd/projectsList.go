package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/juju/ansiterm"
	"github.com/spf13/cobra"
)

var projectsListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Show all projects you're part of",
	Run: func(cmd *cobra.Command, args []string) {
		projects := getProjects()
		currentProject := getCurrentProject()

		writer := ansiterm.NewTabWriter(os.Stdout, 0, 8, 1, '\t', 0)
		faint := color.New(color.Faint).SprintfFunc()
		fmt.Println()
		fmt.Fprintf(writer, "  %s\t%s\n", faint("name"), faint("id"))
		for _, p := range projects {
			name := "  " + p.Name
			if p.Id == currentProject.Id {
				name = "✔ " + p.Name
			}
			fmt.Fprintf(writer, "%-*.*s\t%s\n", 12, 24, name, p.Id)
		}
		writer.Flush()
		fmt.Println()
	},
}

func init() {
	projectsCmd.AddCommand(projectsListCmd)
}
