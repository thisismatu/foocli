package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/juju/ansiterm"
	"github.com/spf13/cobra"
)

var modelsListCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "List available models",
	Run: func(cmd *cobra.Command, args []string) {
		loading("Fetching models", 1)

		currentProject := getCurrentProject()
		models := getModels(currentProject.Id, true)

		fmt.Printf("Models in %s\n", color.CyanString(currentProject.Name))
		fmt.Printf("To see information about a model, run %s\n\n", color.CyanString("`foo model info <model_id>`"))

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
	},
}

func init() {
	modelsCmd.AddCommand(modelsListCmd)
}
