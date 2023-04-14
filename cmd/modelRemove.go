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
		if len(args) == 1 {
			model, err := getModel(args[0])
			if err != nil {
				logError(err)
			}

			handleRemoveModel(model)
		} else {
			currentProject := getCurrentProject()

			selectTemplates := &promptui.SelectTemplates{
				Active:   "▸ {{ .Name | underline }} {{ .Id | faint }}",
				Inactive: "  {{ .Name }} {{ .Id | faint }}",
			}

			adaptedModels := getAdaptedModels(currentProject.Id)
			if len(adaptedModels) == 0 {
				fmt.Println("No models to delete")
				os.Exit(0)
			}

			promptSelect := &promptui.Select{
				Label:        "Select model to delete",
				Items:        adaptedModels,
				Templates:    selectTemplates,
				Stdout:       noBellStdout,
				HideSelected: true,
			}

			i, _, err := promptSelect.Run()
			if err != nil {
				fmt.Println("Cancelled")
				os.Exit(0)
			}

			promptConfirm := promptui.Prompt{
				Label:       fmt.Sprintf("Delete model %s (%s)", adaptedModels[i].Name, adaptedModels[i].Id),
				IsConfirm:   true,
				HideEntered: true,
			}

			_, err = promptConfirm.Run()
			if err != nil {
				fmt.Println("Cancelled")
				os.Exit(0)
			}

			handleRemoveModel(adaptedModels[i])
		}
	},
}

func handleRemoveModel(m Model) {
	removeModel(m)
	logSuccess(fmt.Sprintf("Model %s (%s) was deleted", m.Name, m.Id))
}

func init() {}
