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
		currentProject := getCurrentProject()
		apps := getApplications()

		cyan := color.New(color.FgCyan).SprintFunc()
		if len(apps) == 0 {
			fmt.Printf("No applications in %s\n", cyan(currentProject.Name))
			fmt.Printf("To create an application run %s\n", cyan("foo create [name]"))
			os.Exit(0)
		}
		fmt.Printf("Applications in %s\n\n", cyan(currentProject.Name))

		writer := ansiterm.NewTabWriter(os.Stdout, 0, 8, 1, '\t', 0)
		faint := color.New(color.Faint).SprintfFunc()
		fmt.Fprintf(writer, "  %s\t%s\t%s\t%s\n", faint("name"), faint("id"), faint("status"), faint("deployed"))
		for _, a := range apps {
			date := ""
			if a.Deployed != "" {
				d, _ := time.Parse("2006-01-02 15:04:05 +0000 UTC", a.Deployed)
				date = d.Format("2006-01-02 15:04")
			}
			sc := color.New(statusColor(a.Status)).SprintFunc()
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%s %s\t%s\n", 12, 24, a.Name, a.Id, sc("●"), a.Status, date)
		}
		writer.Flush()
		fmt.Println()
	},
}

func statusColor(status string) color.Attribute {
	c := color.Faint
	switch s := status; s {
	case "Ready":
		c = color.FgGreen
	case "Training":
		c = color.FgYellow
	case "Queued":
		c = color.FgYellow
	case "Failed":
		c = color.FgRed
	default:
		c = color.Faint
	}
	return c
}

func init() {
	rootCmd.AddCommand(listCmd)
}
