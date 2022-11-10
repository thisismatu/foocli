package cmd

import (
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var projectsCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"add"},
	Short:   "Create a new project",

	Run: func(cmd *cobra.Command, args []string) {
		validate := func(input string) error {
			if input == "" {
				return errors.New("invalid name")
			}
			return nil
		}

		templates := &promptui.PromptTemplates{
			Prompt:  "{{ . }} ",
			Invalid: "{{ `?` | blue }} {{ . }}: ",
		}

		prompt := promptui.Prompt{
			Label:     "Project name",
			Templates: templates,
			Validate:  validate,
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		id := uuid.New()
		newProject := Project{Name: result, Id: id.String()}
		addProject(newProject)
		setCurrentProject(newProject.Id)

		name := color.CyanString(result)
		fmt.Printf("Project %s created and set to current project\n", name)
	},
}

func init() {
	projectsCmd.AddCommand(projectsCreateCmd)
}
