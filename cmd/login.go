package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/manifoldco/promptui"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

type Provider struct {
	Name string
	Url  string
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log into your account or create a new one ",
	Run: func(cmd *cobra.Command, args []string) {
		providers := []Provider{
			{Name: "Google", Url: "https://google.com"},
			{Name: "GitHub", Url: "https://github.com"},
		}

		templates := &promptui.SelectTemplates{
			Active:   "▸ {{ .Name | underline }}",
			Inactive: "  {{ .Name }}",
		}

		prompt := &promptui.Select{
			Label:        "Log in using",
			Items:        providers,
			Templates:    templates,
			Stdout:       noBellStdout,
			HideSelected: true,
		}

		i, _, err := prompt.Run()
		if err != nil {
			fmt.Println("No changes made")
			os.Exit(0)
		}

		if providers[i].Url != "" {
			fmt.Printf("Logging in using %s\n", providers[i].Name)
			browser.OpenURL(providers[i].Url)
			fmt.Println()
			fmt.Printf("Visit the following URL if your browser doesn't automatically open: %s/auth/%s\n", providers[i].Url, uuid.New())
			fmt.Println()
			msg := fmt.Sprintf("Waiting for %s authentication to be completed", providers[i].Name)
			loading(msg, 5)
			fmt.Printf("%s You are now logged in\n", color.GreenString("Success!"))
		} else {
			fmt.Println("No changes made")
		}
	},
}

func init() {}
