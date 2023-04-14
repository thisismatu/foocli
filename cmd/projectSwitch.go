package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var projectSwitchCmd = &cobra.Command{
	Use:     "switch [id]",
	Aliases: []string{"use"},
	Short:   "Switch to a different project",
	Args:    cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		currentProject := getCurrentProject()
		projects := getProjects()

		if len(args) == 1 {
			nextProject := Project{}
			for _, p := range projects {
				if p.Id == args[0] {
					nextProject = p
				}
			}

			handleSwitchProject(nextProject, currentProject.Id)
		} else {
			loading("Fetching projects", 1)

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
				fmt.Println("Cancelled")
				os.Exit(0)
			}

			handleSwitchProject(projects[i], currentProject.Id)
		}
	},
}

func handleSwitchProject(p Project, pId string) {
	if p.Id != pId {
		setCurrentProject(p.Id)
		logSuccess(fmt.Sprintf("Switched to project %s (%s)", p.Name, p.Id))
	} else {
		fmt.Println("No changes made")
	}
}

func init() {}
