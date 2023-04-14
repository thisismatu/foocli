package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var projectInviteCmd = &cobra.Command{
	Use:   "invite",
	Short: "Invite a user to the current project",
	Run: func(cmd *cobra.Command, args []string) {
		currentProject := getCurrentProject()
		id := uuid.New()
		fmt.Printf("Share the link with the user you want to invite to project %s\n", color.CyanString(currentProject.Name))
		fmt.Println()
		color.Cyan("https://foo.bar/invite/%s", id)
		fmt.Println()
		fmt.Println("Invite links are one-time links that require the user to have a Foo account. If you need to invite several users, generate a link for each of them.")
	},
}

func init() {}
