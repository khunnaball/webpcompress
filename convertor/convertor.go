package convertor

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Convert takes an input image file and converts it to WebP format,
// saving the result to the specified output directory at the given quality.
// It returns the output file path and any error encountered.
func Convert(inputPath string, outputDir string, quality int) (string, error) {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", fmt.Errorf("could not create output directory: %w", err)
	}

	outputPath := OutputPath(inputPath, outputDir)

	// Build the cwebp command
	cmd := exec.Command("cwebp", "-q", fmt.Sprintf("%d", quality), inputPath, "-o", outputPath)

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("cwebp failed for %s: %w", inputPath, err)
	}

	return outputPath, nil
}

// IsImage reports whether the given file path has a supported image extension.
func IsImage(path string) bool {
	supported := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".tiff": true,
		".bmp":  true,
		".webp": true,
	}

	ext := strings.ToLower(filepath.Ext(path))
	return supported[ext]
}

// FileExists reports whether a file exists at the given path.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

// OutputPath returns the expected output path for a given input file and output directory.
func OutputPath(inputPath string, outputDir string) string {
	filename := filepath.Base(inputPath)
	extension := filepath.Ext(filename)
	baseName := strings.TrimSuffix(filename, extension)
	return filepath.Join(outputDir, baseName+".webp")
}
