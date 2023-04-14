package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/juju/ansiterm"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var projectAddCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"create"},
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
			Label:       "Project name",
			Templates:   templates,
			Validate:    validate,
			HideEntered: true,
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Println("Cancelled")
			os.Exit(0)
		}

		id := uuid.New()
		newProject := Project{Name: result, Id: id.String(), PaymentPlan: "Starter", UserCount: 1}
		addProject(newProject)
		setCurrentProject(newProject.Id)

		writer := ansiterm.NewTabWriter(os.Stdout, 0, 8, 2, '\t', 0)
		fmt.Println()
		fmt.Fprintf(writer, "  %s\t%s\n", faint("Name"), newProject.Name)
		fmt.Fprintf(writer, "  %s\t%s\n", faint("ID"), newProject.Id)
		writer.Flush()
		fmt.Println()
		logSuccess("Project created and set as the current project")
	},
}

func init() {}
