package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/juju/ansiterm"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List applications in the current project",
	Run: func(cmd *cobra.Command, args []string) {
		loading("Fetching applications", 1)

		currentProject := getCurrentProject()
		apps := getApplications()

		if len(apps) == 0 {
			fmt.Printf("No applications in %s\n", color.CyanString(currentProject.Name))
			fmt.Printf("To create an application run %s\n", color.CyanString("`foo create [name]`"))
			os.Exit(0)
		}
		fmt.Printf("Applications in %s\n\n", color.CyanString(currentProject.Name))

		writer := ansiterm.NewTabWriter(os.Stdout, 0, 8, 1, '\t', 0)
		writer.SetStyle(ansiterm.Style(2))
		fmt.Fprintf(writer, "  %s\t%s\t%s\t%s\t%s\n", "name", "id", "language", "status", "deployed")
		writer.Reset()
		for _, a := range apps {
			date := ""
			if a.Deployed != "" {
				d, _ := time.Parse("2006-01-02 15:04:05 +0000 UTC", a.Deployed)
				date = d.Format("2006-01-02 15:04")
			}
			sc := color.New(statusColor(a.Status)).SprintFunc()
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%s\t%s %s\t%s\n", 8, 32, a.Name, a.Id, a.Language, sc("●"), a.Status, date)
		}
		writer.Flush()
		fmt.Println()
	},
}

func statusColor(status string) color.Attribute {
	switch status {
	case "Ready":
		return color.FgGreen
	case "Training":
		return color.FgYellow
	case "Queued":
		return color.FgYellow
	case "Failed":
		return color.FgRed
	default:
		return color.Faint
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
