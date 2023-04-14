package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var modelDeployCmd = &cobra.Command{
	Use:   "deploy [id]",
	Short: "Perform a deployment",
	Long:  "Uploads the contents of the directory to the cloud for validation and training.\nOnce the training is successfully completed, a new version of the adapted model is deployed.",
	Example: `  Deploy the current directory
  $ foo model deploy my-model-id

  Deploy a custom path
  $ foo model deploy my-model-id /path/to/dir`,
	Args: cobra.RangeArgs(0, 2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		model, err := getModel(args[0])
		if err != nil {
			logError(err)
		}

		dir, _ := os.Getwd()
		if len(args) == 2 {
			dir = args[1]
		}
		loading("Deploying changes in "+dir, 1)

		wait, _ := cmd.Flags().GetBool("wait")
		if wait {
			loading("Training model, previous training took 10 seconds", 10)
		}

		logSuccess("Changes deployed")
		fmt.Printf("Inspect the deployment at http://dashboard.foo.com/model/%s\n", model.Id)
	},
}

func init() {
	modelDeployCmd.Flags().BoolP("wait", "w", false, "Wait for training to be finished")
}
