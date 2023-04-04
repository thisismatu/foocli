/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var modelsCreateCmd = &cobra.Command{
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
			Label:     "Model name",
			Templates: promptTemplates,
			Validate:  validate,
		}

		name, err := promptInput.Run()
		if err != nil {
			logError(err)
			return
		}

		selectTemplates := &promptui.SelectTemplates{
			Active:   "{{ `▸` | cyan }} {{ .Name | cyan }}",
			Inactive: "  {{ .Name }}",
			Selected: "{{ `Base model:` | faint }} {{ .Name }}",
		}

		models := getBaseModels()
		promptSelect := &promptui.Select{
			Label:     "Select base model",
			Items:     models,
			Templates: selectTemplates,
			Stdout:    noBellStdout,
		}

		mid, _, err := promptSelect.Run()
		if err != nil {
			logError(err)
			return
		}

		currentProject := getCurrentProject()
		id := uuid.New()
		newModel := Model{Name: name, Language: models[mid].Language, Id: id.String(), ProjectId: currentProject.Id, Status: "Ready"}
		addModel(currentProject.Id, newModel)

		faint := color.New(color.Faint).SprintFunc()
		fmt.Printf("%s %s\n", faint("Model ID:"), id.String())
		fmt.Printf("%s %s\n", faint("Model language:"), models[mid].Language)
		fmt.Println()
		fmt.Println("Adapted model created")
	},
}

func init() {
	modelsCmd.AddCommand(modelsCreateCmd)
}
