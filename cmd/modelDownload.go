package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var modelDownloadCmd = &cobra.Command{
	Use:   "download [id]",
	Short: "Download training data and config files",
	Long:  `Downloads the latest adapted model training data and config files from the cloud`,
	Example: fmtExample("Download to current directory", "foo download my-model-id", false) +
		fmtExample("Download to custom path", "foo download my-model-id /path/to/dir", true),
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		_, err := getModel(args[0])
		if err != nil {
			logError(err)
		}

		dir, _ := os.Getwd()
		if len(args) == 1 {
			dir = args[0]
		}

		loading("Downloading training data to"+dir, 2)
		logSuccess("Wrote 3 files to " + dir)
	},
}

func init() {}
