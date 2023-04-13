package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var modelRemoveCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove"},
	Short:   "Delete adapted model",
	Run: func(cmd *cobra.Command, args []string) {
		currentProject := getCurrentProject()

		selectTemplates := &promptui.SelectTemplates{
			Active:   "{{ `▸` | cyan }} {{ .Name | cyan }} {{ .Id | faint }}",
			Inactive: "  {{ .Name }} {{ .Id | faint }}",
			Selected: "{{ `Model:` | faint }} {{ .Name }} {{ .Id }}",
		}

		models := getModels(currentProject.Id, false)
		if len(models) == 0 {
			fmt.Println("No models to delete")
			os.Exit(0)
		}

		promptSelect := &promptui.Select{
			Label:     "Select model to delete",
			Items:     models,
			Templates: selectTemplates,
			Stdout:    noBellStdout,
		}

		mid, _, err := promptSelect.Run()

		if err != nil {
			os.Exit(0)
		}

		promptConfirm := promptui.Prompt{
			Label:     "Delete model",
			IsConfirm: true,
		}

		_, err = promptConfirm.Run()

		if err != nil {
			os.Exit(0)
		}

		removeModel(currentProject.Id, models[mid].Id)
		fmt.Printf("%s %s was deleted\n", models[mid].Name, models[mid].Id)
	},
}

func init() {
	modelCmd.AddCommand(modelRemoveCmd)
}
