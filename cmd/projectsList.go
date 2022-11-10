/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/juju/ansiterm"
	"github.com/spf13/cobra"
)

// projectsListCmd represents the projectList command
var projectsListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Show all projects you're part of",
	Run: func(cmd *cobra.Command, args []string) {
		projects := getProjects()
		currentProject := getCurrentProject()

		fmt.Print("Your projects\n\n")

		writer := ansiterm.NewTabWriter(os.Stdout, 0, 8, 1, '\t', 0)
		gray := color.New(color.Faint).SprintfFunc()
		fmt.Fprintf(writer, "  %s\t%s\n", gray("name"), gray("id"))
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectsListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectsListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
