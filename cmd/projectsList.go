package cmd

import (
	"fmt"
	"os"

	"github.com/juju/ansiterm"
	"github.com/spf13/cobra"
)

var projectsListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Show all projects you're part of",
	Run: func(cmd *cobra.Command, args []string) {
		loading("Fetching projects", 1)

		currentProject := getCurrentProject()
		projects := getProjects()

		writer := ansiterm.NewTabWriter(os.Stdout, 0, 8, 2, '\t', 0)
		writer.SetStyle(ansiterm.Style(2))
		fmt.Println()
		fmt.Fprintf(writer, "  %s\t%s\n", "Name", "Project ID")
		writer.Reset()
		for _, p := range projects {
			name := "  " + p.Name
			if p.Id == currentProject.Id {
				name = "✔ " + p.Name
			}
			fmt.Fprintf(writer, "%-*.*s\t%s\n", 8, 32, name, p.Id)
		}
		writer.Flush()
		fmt.Println()
	},
}

func init() {
	projectsCmd.AddCommand(projectsListCmd)
}
