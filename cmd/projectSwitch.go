package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var projectSwitchCmd = &cobra.Command{
	Use:     "switch [project]",
	Aliases: []string{"use"},
	Short:   "Switch to a different project",
	Run: func(cmd *cobra.Command, args []string) {
		loading("Fetching projects", 1)

		currentProject := getCurrentProject()
		projects := getProjects()
		for i, p := range projects {
			if p.Id == currentProject.Id {
				projects[i].Name = p.Name + " (current)"
			}
		}

		templates := &promptui.SelectTemplates{
			Active:   "▸ {{ .Name | underline }} {{ .Id | faint }}",
			Inactive: "  {{ .Name }} {{ .Id | faint }}",
		}

		prompt := &promptui.Select{
			Label:        "Select project",
			Items:        projects,
			Templates:    templates,
			Stdout:       noBellStdout,
			HideSelected: true,
		}

		i, _, err := prompt.Run()
		if err != nil {
			fmt.Println("No changes made")
			os.Exit(0)
		}

		if projects[i].Id != "" && projects[i].Id != currentProject.Id {
			setCurrentProject(projects[i].Id)
			fmt.Printf("%s Switched to project %s (%s)\n", color.GreenString("Success!"), projects[i].Name, projects[i].Id)
		} else {
			fmt.Println("No changes made")
		}
	},
}

func init() {
	projectCmd.AddCommand(projectSwitchCmd)
}
