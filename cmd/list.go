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

type app struct {
	Name     string
	Id       string
	Status   string
	Deployed string
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List applications in the current project",
	Run: func(cmd *cobra.Command, args []string) {
		apps := []app{
			{Name: "Prod app", Id: "523f7849-d68d-46bc-8b6b-9522afa04557", Status: "Ready", Deployed: "2022-11-09 12:31:21.235527 +0000 UTC"},
			{Name: "Dev app", Id: "3e618ca8-be49-43b1-9d00-6083a8e48cf1", Status: "Failed", Deployed: ""},
			{Name: "Testing", Id: "2f567022-798c-47ea-9047-adf30055d692", Status: "Ready", Deployed: "2022-11-08 11:37:19.093058 +0000 UTC"},
			{Name: "Foobar", Id: "a77cf94c-586d-4409-a342-e58ed6b163ad", Status: "Ready", Deployed: "2022-11-08 08:26:38.292301 +0000 UTC"},
		}

		fmt.Print("Applications in My project\n\n")

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
