package fileutils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// GetFilePaths returns a list of file paths in the provided root directory and its
// subdirectories that match any of the provided patterns.
func GetFilePaths(includePatterns []string, rootDir string) ([]string, error) {
	filePaths := []string{}
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		for _, pattern := range includePatterns {
			matched, err := filepath.Match(pattern, filepath.Base(path))
			if err != nil {
				return err
			}

			if matched {
				// add the file path to the list of file paths, but remove the root directory
				// from the path and / from the beginning of the path, but only if the rootDir is not "."
				if rootDir != "." {
					filePaths = append(filePaths, strings.TrimPrefix(path, rootDir)[1:])
				} else {
					filePaths = append(filePaths, path)
				}
				return nil
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return filePaths, err
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
			if matched, _ := filepath.Match(pattern, file); matched {
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
