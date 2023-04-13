/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var modelAddCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"create"},
	Short:   "Create a new adapted model",
	Run: func(cmd *cobra.Command, args []string) {
		validate := func(input string) error {
			if input == "" {
				return errors.New("invalid name")
			}
			return nil
		}

		promptTemplates := &promptui.PromptTemplates{
			Prompt:  "{{ . }} ",
			Invalid: "{{ `?` | blue }} {{ . }}: ",
		}

		promptInput := promptui.Prompt{
			Label:       "Model name",
			Templates:   promptTemplates,
			Validate:    validate,
			HideEntered: true,
		}

		name, err := promptInput.Run()
		if err != nil {
			fmt.Println("No changes made")
			os.Exit(0)
		}

		selectTemplates := &promptui.SelectTemplates{
			Active:   "▸ {{ .Name | underline }} {{ .Id | faint }}",
			Inactive: "  {{ .Name }} {{ .Id | faint }}",
		}

		baseModels := getBaseModels()
		promptSelect := &promptui.Select{
			Label:        "Select base model",
			Items:        baseModels,
			Templates:    selectTemplates,
			Stdout:       noBellStdout,
			HideSelected: true,
		}

		i, _, err := promptSelect.Run()
		if err != nil {
			fmt.Println("No changes made")
			os.Exit(0)
		}

		currentProject := getCurrentProject()
		id := uuid.New()
		newModel := Model{Name: name, Language: baseModels[i].Language, Id: id.String(), ProjectId: currentProject.Id, Status: "Ready", BaseModel: baseModels[i].Id}
		addModel(currentProject.Id, newModel)
		printModelInfo(newModel)
		logSuccess("Adapted model created")
	},
}

func init() {
	modelCmd.AddCommand(modelAddCmd)
}
