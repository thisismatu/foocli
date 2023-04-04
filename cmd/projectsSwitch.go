package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var projectsSwitchCmd = &cobra.Command{
	Use:     "switch",
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
		projects = append(projects, Project{Name: "Cancel"})

		templates := &promptui.SelectTemplates{
			Active:   "{{ `▸` | cyan }} {{ .Name | cyan }}",
			Inactive: "  {{ .Name }}",
			Selected: "{{ if .Id }}Switched to project {{ .Name | cyan }}{{ else }}No changes made{{ end }}",
		}

		prompt := promptui.Select{
			Label:     "Switch project",
			Items:     projects,
			Templates: templates,
			Stdout:    noBellStdout,
		}

		i, _, err := prompt.Run()
		if err != nil {
			return
		}

		if projects[i].Id != "" {
			setCurrentProject(projects[i].Id)
		}
	},
}

func init() {
	projectsCmd.AddCommand(projectsSwitchCmd)
}
