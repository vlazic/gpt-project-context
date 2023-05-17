package fileutils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GetFilePaths(includePatterns []string) ([]string, error) {
	filePaths := []string{}
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
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
				filePaths = append(filePaths, path)
				return nil
			}
		}

		return nil
	})

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
