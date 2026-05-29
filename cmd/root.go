package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/khunnaball/webpcompress/convertor"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "webpcompress",
	Short: "A CLI tool for converting and compressing images to WebP",
	Long:  "webpcompress converts images to WebP with configurable quality and output options.",
	Run: func(cmd *cobra.Command, args []string) {
		quality, _ := cmd.Flags().GetInt("quality")
		output, _ := cmd.Flags().GetString("output")
		dir, _ := cmd.Flags().GetString("dir")
		force, _ := cmd.Flags().GetBool("force")

		var files []string

		if quality < 1 || quality > 100 {
			fmt.Println("Error: quality must be between 1 and 100")
			os.Exit(1)
		}

		if dir != "" {
			matches, err := filepath.Glob(filepath.Join(dir, "*"))
			if err != nil {
				fmt.Println("Error reading directory:", err)
			}
			files = matches
		} else {
			if len(args) == 0 {
				fmt.Println("Error: no input files provided.")
				cmd.Help()
				os.Exit(1)
			}
			files = args
		}

		var converted, skipped, errCount int

		// Loop over files and convert them
		for _, file := range files {
			if !convertor.IsImage(file) {
				continue
			}

			outputPath := convertor.OutputPath(file, output)

			if convertor.FileExists(outputPath) && !force {
				fmt.Printf("Skipping %s: already exists (use --force to overwrite)\n", file)
				skipped++
				continue
			}

			outputPath, err := convertor.Convert(file, output, quality)
			// Log error and continue with next file
			if err != nil {
				fmt.Printf("Error converting %s: %s\n", file, err)
				errCount++
				continue
			}
			fmt.Printf("Converted: %s -> %s\n", file, outputPath)
			converted++
		}

		fmt.Printf("Done: %d converted, %d skipped, %d error(s)\n", converted, skipped, errCount)
	},
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntP("quality", "q", 80, "Webp quality (1-100)")
	rootCmd.Flags().StringP("output", "o", "./output", "Output directory (optional)")
	rootCmd.Flags().StringP("dir", "d", "", "Directory of images to convert (optional)")
	rootCmd.Flags().BoolP("force", "f", false, "Overwrite existing files")
}
