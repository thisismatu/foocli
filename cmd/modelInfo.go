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
	Use:   "info [model]",
	Short: "Display information about a model",
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
		printModelInfo(model)

		if model.ProjectId != "all" {
			writer := ansiterm.NewTabWriter(os.Stdout, 0, 8, 2, '\t', 0)
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
		}
	},
}

func init() {
	modelCmd.AddCommand(modelInfoCmd)
}
