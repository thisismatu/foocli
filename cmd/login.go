package cmd

import (
	"fmt"

	"github.com/fatih/color"
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
	Short: "Log into the Foo platfrom",
	Run: func(cmd *cobra.Command, args []string) {
		providers := []Provider{
			{Name: "Google", Url: "https://google.com"},
			{Name: "GitHub", Url: "https://github.com"},
			{Name: "Cancel"},
		}

		templates := &promptui.SelectTemplates{
			Active:   "{{ `▸` | cyan }} {{ .Name | cyan }}",
			Inactive: "  {{ .Name }}",
			Selected: "Log in to Foo {{ .Name | cyan }}",
		}

		prompt := promptui.Select{
			Label:     "Log in to Foo",
			Items:     providers,
			Templates: templates,
			Stdout:    noBellStdout,
		}

		i, _, err := prompt.Run()
		if err != nil {
			return
		}

		if providers[i].Url != "" {
			browser.OpenURL(providers[i].Url)
			fmt.Printf("Visit the following URL if your browser doesn't automatically open: %s\n", providers[i].Url)
			msg := fmt.Sprintf("Waiting for %s authentication to be completed", providers[i].Name)
			loading(msg, 5)
			success := color.CyanString("Success!")
			fmt.Printf("%s You are now logged in.\n", success)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
