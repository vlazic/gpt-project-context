package fileutils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
)

// GetFilePaths returns a list of file paths in the provided root directory and its
// subdirectories that match any of the provided patterns.
// GetFilePaths returns a list of file paths in the provided root directory and its
// subdirectories that match any of the provided patterns.
func GetFilePaths(includePatterns []string, rootDir string) ([]string, error) {
	filePaths := []string{}

	for _, pattern := range includePatterns {
		fs := os.DirFS(rootDir)
		matches, err := doublestar.Glob(fs, pattern)
		if err != nil {
			return nil, err
		}

		// add the file paths to the list of file paths, but remove the root directory
		// from the path and / from the beginning of the path, but only if the rootDir is not "."
		for _, match := range matches {
			info, err := os.Stat(filepath.Join(rootDir, match))
			if err != nil {
				return nil, err
			}
			if !info.IsDir() {
				filePaths = append(filePaths, match)
			}
		}
	}

	return filePaths, nil
}

func FilterFiles(filePaths, excludePatterns []string) []string {
	filteredFiles := []string{}
	defaultExclusions := []string{".git"}
	for _, file := range filePaths {
		exclude := false

		// Check for default exclusions
		for _, defaultExclusion := range defaultExclusions {
			if strings.Contains(file, defaultExclusion) {
				exclude = true
				break
			}
		}

		// Check for user defined exclusions
		for _, pattern := range excludePatterns {
			matched, err := doublestar.Match(pattern, file)
			if err != nil {
				panic(err)
			}
			if matched {
				exclude = true
				break
			}
		}
		if !exclude {
			filteredFiles = append(filteredFiles, file)
		}
	}
	return filteredFiles
}

func CreateOutput(filteredFiles []string, prompt string) (string, error) {
	output := prompt + "\n\n"
	for _, file := range filteredFiles {
		output += "---\n"
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return "", err
		}
		output += "# FILE: " + file + "\n\n "
		output += string(content) + "\n"
	}
	return output, nil
}

func WriteToFile(filename string, data string) error {
	return ioutil.WriteFile(filename, []byte(data), 0644)
}
