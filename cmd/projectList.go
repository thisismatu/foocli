/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// projectListCmd represents the projectList command
var projectListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Show all projects you're part of",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Your projects\n\n")

		writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
		header := color.New(color.Faint).FprintfFunc()
		header(writer, "  name\tid\n")
		for _, p := range projects {
			selected := " "
			if p.Id == getCurrentProject() {
				selected = "✔"
			}
			row := color.New(color.FgWhite).FprintfFunc()
			row(writer, "%s %s\t%s\n", selected, p.Name, p.Id)
		}
		writer.Flush()
		fmt.Println()
	},
}

func init() {
	projectsCmd.AddCommand(projectListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
