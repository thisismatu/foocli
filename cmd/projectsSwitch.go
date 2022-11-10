/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// projectsSwitchCmd represents the projectSwitch command
var projectsSwitchCmd = &cobra.Command{
	Use:     "switch",
	Aliases: []string{"use"},
	Short:   "Switch to a different project",
	Run: func(cmd *cobra.Command, args []string) {
		projects := getProjects()
		projects = append(projects, Project{Name: "Cancel"})

		templates := &promptui.SelectTemplates{
			Active:   "{{ `▸` | cyan }} {{ .Name | cyan }}",
			Inactive: "  {{ .Name }}",
			Selected: "Switched to project {{ .Name | cyan }}",
		}

		prompt := promptui.Select{
			Label:     "Switch project",
			Items:     projects,
			Templates: templates,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectsSwitchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectsSwitchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
