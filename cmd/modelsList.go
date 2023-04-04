package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/juju/ansiterm"
	"github.com/spf13/cobra"
)

var modelsListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List available models",
	Args:    cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		mid, err := cmd.Flags().GetString("model")
		if err != nil {
			log.Fatalf("Missing model ID: %s", err)
		}

		if mid == "" && len(args) > 0 {
			mid = args[0]
		}

		if mid == "" {
			loading("Fetching models", 1)

			currentProject := getCurrentProject()
			models := getModels(currentProject.Id)

			fmt.Printf("Models in %s\n", color.CyanString(currentProject.Name))
			fmt.Printf("To list deployments for a model, run %s\n\n", color.CyanString("foo models list <model_id>"))

			writer := ansiterm.NewTabWriter(os.Stdout, 0, 8, 2, '\t', 0)
			writer.SetStyle(ansiterm.Style(2))
			fmt.Fprintf(writer, "  %s\t%s\t%s\t%s\n", "Name", "Language", "Model ID", "Status")
			writer.Reset()
			for _, m := range models {
				sc := color.New(statusColor(m.Status)).SprintFunc()
				fmt.Fprintf(writer, "  %-*.*s\t%s\t%s\t%s %s\n", 8, 32, m.Name, m.Language, m.Id, sc("●"), m.Status)
			}
			writer.Flush()
			fmt.Println()
		} else {
			loading("Fetching deployments", 1)

			currentProject := getCurrentProject()
			deployments := getDeployments(mid)
			model := getModel(mid)

			fmt.Printf("Deployments for %s in %s\n\n", color.CyanString(model.Name), color.CyanString(currentProject.Name))

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
	modelsCmd.AddCommand(modelsListCmd)
	modelsListCmd.Flags().StringP("model", "m", "", "List deployments for a specific model")
}
