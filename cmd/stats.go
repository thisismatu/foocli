package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/juju/ansiterm"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats [id]",
	Short: "Get API usage statistics",
	Long:  `Get API usage statistics for the current project or a specific model`,
	Example: fmtExample("Project API usage", "foo stats", false) +
		fmtExample("Project API usage for a custom time range", "foo stats --start 2023-01-01 --end --2023-03-31", false) +
		fmtExample("Model API usage", "foo stats my-model-id", true),
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 1 {
			model, err := getModel(args[0])
			if err != nil {
				logError(err)
			}
			loading("Loading model usage statistics", 2)
			fmt.Printf("API usage for model %s\n", color.CyanString(model.Name))

			writer := ansiterm.NewTabWriter(os.Stdout, 0, 8, 2, '\t', 0)
			fmt.Println()
			fmt.Fprintf(writer, "  %s\t%s\n", faint("Start time"), "2023-04-04")
			fmt.Fprintf(writer, "  %s\t%s\n", faint("End time"), "2023-04-18")
			fmt.Fprintf(writer, "  %s\t%s\n", faint("Audio sent"), "8 h 41 min 48 s")
			fmt.Fprintf(writer, "  %s\t%s\n", faint("Utterances"), "5124")
			fmt.Fprintf(writer, "  %s\t%s\n", faint("Requests"), "36")
			fmt.Fprintf(writer, "  %s\t%s\n", faint("Annotated audio"), "0")
			writer.Flush()
			fmt.Println()

			writer.SetStyle(ansiterm.Style(2))
			fmt.Fprintf(writer, "  %s\t%s\t%s\t%s\t%s\n", "Date", "Audio sent", "Utterances", "Requests", "Annotated audio")
			writer.Reset()
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%d\t%d\t%d\n", 8, 32, "2023-04-18", "1h 41 min 28 s", 112, 12, 0)
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%d\t%d\t%d\n", 8, 32, "2023-04-17", "41 min 58 s", 54, 8, 0)
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%d\t%d\t%d\n", 8, 32, "2023-04-16", "12 min 32 s", 11, 4, 0)
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%d\t%d\t%d\n", 8, 32, "2023-04-15", "1h 41 min 28 s", 112, 12, 0)
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%d\t%d\t%d\n", 8, 32, "2023-04-14", "41 min 58 s", 54, 8, 0)
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%d\t%d\t%d\n", 8, 32, "2023-04-13", "12 min 32 s", 11, 4, 0)
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%d\t%d\t%d\n", 8, 32, "2023-04-12", "1h 41 min 28 s", 112, 12, 0)
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%d\t%d\t%d\n", 8, 32, "2023-04-11", "41 min 58 s", 54, 8, 0)
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%d\t%d\t%d\n", 8, 32, "2023-04-10", "12 min 32 s", 11, 4, 0)
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%d\t%d\t%d\n", 8, 32, "2023-04-09", "1h 41 min 28 s", 112, 12, 0)
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%d\t%d\t%d\n", 8, 32, "2023-04-08", "41 min 58 s", 54, 8, 0)
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%d\t%d\t%d\n", 8, 32, "2023-04-07", "12 min 32 s", 11, 4, 0)
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%d\t%d\t%d\n", 8, 32, "2023-04-06", "1h 41 min 28 s", 112, 12, 0)
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%d\t%d\t%d\n", 8, 32, "2023-04-05", "41 min 58 s", 54, 8, 0)
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%d\t%d\t%d\n", 8, 32, "2023-04-04", "12 min 32 s", 11, 4, 0)
			writer.Flush()
			fmt.Println()
		} else {
			proj := getCurrentProject()
			loading("Loading project usage statistics", 2)
			fmt.Printf("API usage for project %s\n", color.CyanString(proj.Name))

			writer := ansiterm.NewTabWriter(os.Stdout, 0, 8, 2, '\t', 0)
			fmt.Println()
			fmt.Fprintf(writer, "  %s\t%s\n", faint("Start time"), "2023-04-04")
			fmt.Fprintf(writer, "  %s\t%s\n", faint("End time"), "2023-04-18")
			fmt.Fprintf(writer, "  %s\t%s\n", faint("Audio sent"), "26 h 44 min 20 s")
			fmt.Fprintf(writer, "  %s\t%s\n", faint("Utterances"), "12526")
			fmt.Fprintf(writer, "  %s\t%s\n", faint("Requests"), "123")
			fmt.Fprintf(writer, "  %s\t%s\n", faint("Annotated audio"), "0")
			writer.Flush()
			fmt.Println()

			writer.SetStyle(ansiterm.Style(2))
			fmt.Fprintf(writer, "  %s\t%s\t%s\t%s\t%s\n", "Model", "Audio sent", "Utterances", "Requests", "Annotated audio")
			writer.Reset()
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%d\t%d\t%d\n", 8, 32, "Example application 1", "8 h 41 min 48 s", 5124, 36, 0)
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%d\t%d\t%d\n", 8, 32, "Example application 3", "7 h 13 min 15 s", 4175, 24, 0)
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%d\t%d\t%d\n", 8, 32, "Example application 2", "5 h 24 min 9 s", 3984, 18, 0)
			fmt.Fprintf(writer, "  %-*.*s\t%s\t%d\t%d\t%d\n", 8, 32, "Example application 4", "4 h 58 min 24 s", 1243, 9, 0)
			writer.Flush()
			fmt.Println()
		}
	},
}

func init() {
	statsCmd.Flags().String("start", "", "Start date for statistics")
	statsCmd.Flags().String("end", "", "Start date for statistics")
	statsCmd.Flags().String("csv", "", "Print report as CSV")
}
