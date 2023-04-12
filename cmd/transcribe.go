package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var transcribeCmd = &cobra.Command{
	Use:   "transcribe",
	Short: "Transcribe audio files",
	Long:  "Transcribe audio files",
	Example: `  Transcribe a single file

  foo transcribe file.wav

  - If no model is provided, the default model is used

  Transcribe multiple files using a specific model

  foo transcribe files.jsonl <model_id>

  - Specify multiple audio files in a JSON Lines document using the format {"audio":"/path/to/file"}`,
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
}
