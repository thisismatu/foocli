package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/juju/ansiterm"
	"github.com/spf13/cobra"
)

var modelInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display information about a model",
	Long:  "Displays information and deployment history related to a model",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		modelId := args[0]
		deployments := getDeployments(modelId)
		model, err := getModel(modelId)
		if err != nil {
			logError(err)
		}

		loading(fmt.Sprintf("Fetching information for %s", modelId), 1)

		writer := ansiterm.NewTabWriter(os.Stdout, 0, 8, 2, '\t', 0)
		sc := color.New(statusColor(model.Status)).SprintFunc()
		faint := color.New(color.Faint).SprintFunc()

		fmt.Println("General")
		fmt.Println()
		fmt.Fprintf(writer, "  %s\t%s\n", faint("Name"), model.Name)
		fmt.Fprintf(writer, "  %s\t%s\n", faint("ID"), model.Id)
		fmt.Fprintf(writer, "  %s\t%s\n", faint("Language"), model.Language)
		fmt.Fprintf(writer, "  %s\t%s %s\n", faint("Status"), sc("●"), model.Status)
		fmt.Fprintf(writer, "  %s\t%s\n", faint("Description"), "Model description goes here. It should briefly describe the model characteristics.")
		writer.Flush()
		fmt.Println()

		fmt.Println("Deployments")
		fmt.Println()
		writer.SetStyle(ansiterm.Style(2))
		fmt.Fprintf(writer, "  %s\t%s\t%s\t%s\n", "Date", "Deployment", "Status", "Duration")
		writer.Reset()
		for _, d := range deployments {
			date := d.Date.Format(time.Stamp)
			sc := color.New(statusColor(d.Status)).SprintFunc()
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%s %s\t%s\n", 8, 32, date, d.Url, sc("●"), d.Status, d.Duration)
		}
		writer.Flush()
		fmt.Println()
	},
}

func init() {
	modelCmd.AddCommand(modelInfoCmd)
}
