package cmd

import (
	"fmt"

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
		}

		i, _, err := prompt.Run()
		if err != nil {
			return
		}

		if providers[i].Url != "" {
			browser.OpenURL(providers[i].Url)
			fmt.Printf("Please visit the following URL in your browser if it doesn't automatically open:\n%s\n", providers[i].Url)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
