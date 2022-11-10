/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List applications in the current project",
	Run: func(cmd *cobra.Command, args []string) {
		currentProject := getCurrentProject()
		apps := getApplications()

		fmt.Printf("Applications in %s\n\n", currentProject.Name)

		writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
		header := color.New(color.Faint).FprintfFunc()
		header(writer, "  name\tid\tstatus\tdeployed\n")
		for _, a := range apps {
			date := ""
			if a.Deployed != "" {
				d, _ := time.Parse("2006-01-02 15:04:05 +0000 UTC", a.Deployed)
				date = d.Format("2006-01-02 15:04")
			}
			row := color.New(color.FgWhite).FprintfFunc()
			row(writer, "  %s\t%s\t%s\t%s\n", a.Name, a.Id, a.Status, date)
		}
		writer.Flush()
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
