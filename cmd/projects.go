/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

type project struct {
	Name string
	Id   string
}

var projects = []project{
	{Name: "My project", Id: "1a5eedfe-d053-458f-9e2a-02f8878bf220"},
	{Name: "Acme corp", Id: "a0f4652a-b5ca-4e4b-be14-723b359ace72"},
}

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "View and manage projects",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(projectsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// projectsCmd.PersistentFlags().String("crete", "", "Create a new project")
	// projectsCmd.PersistentFlags().String("list", "", "Show all projects you're part of")
	// projectsCmd.PersistentFlags().String("switch", "", "Switch to a different project")
	// projectsCmd.PersistentFlags().String("invite", "", "Invite a user to this project")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
