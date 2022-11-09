/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// projectSwitchCmd represents the projectSwitch command
var projectSwitchCmd = &cobra.Command{
	Use:     "switch",
	Aliases: []string{"use"},
	Short:   "Switch to a different project",
	Run: func(cmd *cobra.Command, args []string) {
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

		setCurrentProject(projects[i].Id)
	},
}

func init() {
	projectsCmd.AddCommand(projectSwitchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectSwitchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectSwitchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
