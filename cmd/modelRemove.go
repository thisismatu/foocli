package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var modelRemoveCmd = &cobra.Command{
	Use:     "rm [model]",
	Aliases: []string{"remove"},
	Short:   "Delete adapted model",
	Run: func(cmd *cobra.Command, args []string) {
		currentProject := getCurrentProject()

		selectTemplates := &promptui.SelectTemplates{
			Active:   "▸ {{ .Name | underline }} {{ .Id | faint }}",
			Inactive: "  {{ .Name }} {{ .Id | faint }}",
		}

		models := getModels(currentProject.Id, false)
		if len(models) == 0 {
			fmt.Println("No models to delete")
			os.Exit(0)
		}

		promptSelect := &promptui.Select{
			Label:        "Select model to delete",
			Items:        models,
			Templates:    selectTemplates,
			Stdout:       noBellStdout,
			HideSelected: true,
		}

		mid, _, err := promptSelect.Run()
		if err != nil {
			fmt.Println("No changes made")
			os.Exit(0)
		}

		promptConfirm := promptui.Prompt{
			Label:       fmt.Sprintf("Delete model %s (%s)", models[mid].Name, models[mid].Id),
			IsConfirm:   true,
			HideEntered: true,
		}

		_, err = promptConfirm.Run()

		if err != nil {
			fmt.Println("No changes made")
			os.Exit(0)
		}

		removeModel(currentProject.Id, models[mid].Id)
		logSuccess(fmt.Sprintf("Model %s (%s) was deleted", models[mid].Name, models[mid].Id))
	},
}

func init() {
	modelCmd.AddCommand(modelRemoveCmd)
}
