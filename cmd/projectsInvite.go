package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var projectsInviteCmd = &cobra.Command{
	Use:   "invite",
	Short: "Invite a user to this project",
	Run: func(cmd *cobra.Command, args []string) {
		currentProject := getCurrentProject()
		id := uuid.New()
		cyan := color.New(color.FgCyan).SprintFunc()
		faint := color.New(color.Faint).PrintlnFunc()
		fmt.Printf("Share this link with the user you want to invite to %s\n\n", cyan(currentProject.Name))
		color.Cyan("  https://foo.bar/invite/%s\n\n", id)
		faint("Invite links are one-time links that require the user to have a Foo account. If you need to invite several users, generate a link for each of them.")
	},
}

func init() {
	projectsCmd.AddCommand(projectsInviteCmd)
}
