package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var transcribeCmd = &cobra.Command{
	Use:   "transcribe [path]",
	Short: "Transcribe audio files",
	Example: fmtExample("Transcribe a single file using the default model", "foo transcribe file.wav", false) +
		fmtExample("Transcribe multiple files using a specific model. Define the files\n  in a JSON Lines document using the format {\"audio\":\"/path/to/file\"}", "foo transcribe files.jsonl small-lowlatency-en", true),
	Args: cobra.RangeArgs(0, 2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		input := args[0]
		absPath, err := filepath.Abs(input)
		if err != nil {
			logError(err)
		}
		file := filepath.Base(absPath)
		fileExt := filepath.Ext(file)[1:]
		validFormats := []string{"mp3", "wav", "flac", "ogg"}
		if !contains(validFormats, fileExt) {
			err := fmt.Errorf("invalid file format, supported formats are %s", strings.Join(validFormats, ", "))
			logError(err)
		}

		modelId := "large-highaccuracy-en"
		if len(args) > 1 {
			modelId = args[1]
		}
		_, err = getModel(modelId)
		if err != nil {
			logError(err)
		}

		loading(fmt.Sprintf("Uploading %s", file), 2)
		loading(fmt.Sprintf("Transcribing %s", file), 5)

		fmt.Println("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam nec augue placerat, mattis mauris at, dapibus quam. Maecenas porta eget arcu id laoreet. Maecenas tincidunt auctor nisi, in rhoncus odio condimentum vitae. In interdum velit orci, non feugiat libero efficitur nec. Maecenas lobortis risus libero, nec accumsan augue lacinia at. Nulla varius arcu eget massa consequat vestibulum. Fusce placerat felis diam, at pulvinar ante accumsan eget.")
	},
}

func init() {
	transcribeCmd.Flags().BoolP("streaming", "s", false, "Use Streaming API instead of Batch API")
}
