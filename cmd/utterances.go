package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var utterancesCmd = &cobra.Command{
	Use:   "utterances",
	Short: "Get a sample of recent utterances",
	Long:  `Fetches a sample of recent utterances`,
	Example: fmtExample("Get recent utterances for the default model", "foo utterances", false) +
		fmtExample("Get recent utterances for a specific model", "foo utterances my-model-id", false) +
		fmtExample("Write utterances to a text file", "foo utterances my-model-id -e > training-data.txt", true),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			model, err := getModel(args[0])
			if err != nil {
				logError(err)
			}
			loading("Fetching utterances", 2)
			fmt.Printf("Utterances for model %s\n", color.CyanString(model.Name))
		} else {
			proj := getCurrentProject()
			loading("Fetching utterances", 2)
			fmt.Printf("Utterances for project %s\n", color.CyanString(proj.Name))
		}

		fmt.Println()
		fmt.Printf("  %s\n", faint("2023-04-16 (1)"))
		fmt.Println("  hi i'm neil degrasse tyson astrophysicist in addition to probing the secrets of the universe i'm also a movie buff today i introduce you to a film that everyone thought was lost forever until a print was recently discovered in a hollywood vault future thirty eight forgotten treasure from nineteen thirty eight it's one of the first collar pictures preceding gone with the wind and the wizard of oz by a year but what interests me most is the science finally a movie that gets time travel right the concepts and technology imagine for the film were so advanced they inspired an entire generation of scientists and engineers and guess what when our hero travels to the distant future when does he land the year two thousand eighteen how do those nineteen thirty eight filmmakers imagine our time what do they get wrong but more importantly what did they get right let's find out we now look back at a film that looks forward to today roll projector")
		fmt.Println()
		fmt.Printf("  %s\n", faint("2023-04-14 (1)"))
		fmt.Println("  hello hello hello do you hear me ahoo ahoy")
		fmt.Println()
		fmt.Printf("  %s\n", faint("2023-04-12 (1)"))
		fmt.Println("  generate sales report for the month of january")
		fmt.Println()
		fmt.Printf("  %s\n", faint("2023-04-11 (2)"))
		fmt.Println("  i actually haven't seen a transcription service since better than this on")
		fmt.Println("  whisper might be better but whispered me more expensive")
		fmt.Println()
	},
}

func init() {
	utterancesCmd.Flags().BoolP("export", "e", false, "Strip dates and padding")
}
