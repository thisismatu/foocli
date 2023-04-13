package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/juju/ansiterm"

	"github.com/fatih/color"
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

		name, err := prompt.Run()
		if err != nil {
			fmt.Println("No changes made")
			os.Exit(0)
		}

		id := uuid.New()
		newProject := Project{Name: name, Id: id.String()}
		addProject(newProject)
		setCurrentProject(newProject.Id)

		writer := ansiterm.NewTabWriter(os.Stdout, 0, 8, 2, '\t', 0)
		faint := color.New(color.Faint).SprintFunc()
		fmt.Println()
		fmt.Fprintf(writer, "  %s\t%s\n", faint("Name"), name)
		fmt.Fprintf(writer, "  %s\t%s\n", faint("ID"), id.String())
		writer.Flush()
		fmt.Println()
		fmt.Printf("%s Project created and set as the current project\n", color.GreenString("Success!"))
	},
}

func init() {
	projectCmd.AddCommand(projectAddCmd)
}
