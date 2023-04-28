package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.design/x/clipboard"
)

var (
	// default prompt at the end of concatenated files
	defaultPrompt = "Here is the context of my current project. Just respond with 'OK' and wait for the instructions:"
)

func getFlags() (include, exclude, prompt, outputFile string, dryRun bool) {
	flag.StringVar(&include, "i", "", "include patterns")
	flag.StringVar(&exclude, "e", "", "exclude patterns")
	flag.StringVar(&prompt, "p", defaultPrompt, "prompt at the beginning")
	flag.StringVar(&outputFile, "o", "", "output file path")
	flag.BoolVar(&dryRun, "n", false, "no action, do not copy or write to clipboard")

	flag.Parse()
	return
}

func getFilePaths(includePatterns []string) ([]string, error) {
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

func filterFiles(filePaths, excludePatterns []string) []string {
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

func confirmAction(filteredFiles []string) {
	fmt.Println("The following files will be copied:")
	for _, file := range filteredFiles {
		fmt.Println(file)
	}
	var response string
	fmt.Print("Continue? [y/N] ")
	fmt.Scanln(&response)
	if response != "y" && response != "Y" {
		os.Exit(1)
	}
}

func createOutput(filteredFiles []string, prompt string) (string, error) {
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

func writeToClipboard(output string) error {
	// Initialize the clipboard
	err := clipboard.Init()
	if err != nil {
		return err
	}

	// If clipboard content is the same as output, do nothing
	origClipData := clipboard.Read(clipboard.FmtText)
	if string(origClipData) == output {
		return nil
	}

	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Watch the clipboard for changes
	clipDataChan := clipboard.Watch(ctx, clipboard.FmtText)

	// Write to clipboard
	clipboard.Write(clipboard.FmtText, []byte(output))

	// Wait for the new data
	var clipData []byte
	select {
	case clipData = <-clipDataChan:
	case <-ctx.Done():
		return fmt.Errorf("timeout waiting for clipboard update")
	}

	// Check if clipboard data is the same as output
	if string(clipData) != output {
		return fmt.Errorf("clipboard content does not match the original output")
	}

	return nil
}

func main() {
	include, exclude, prompt, outputFile, dryRun := getFlags()

	includePatterns := strings.Split(include, ",")
	excludePatterns := strings.Split(exclude, ",")

	filePaths, err := getFilePaths(includePatterns)
	if err != nil {
		fmt.Println("Error while walking the file tree:", err)
		return
	}

	filteredFiles := filterFiles(filePaths, excludePatterns)

	if len(filteredFiles) == 0 {
		fmt.Println("No files found.")
		// sample usage
		fmt.Println("Example usage Go: Usage: context -i '*.go,*.md' -e 'bin/*,*.txt'")
		fmt.Println("Example usage JS: Usage: context -i '*.js,README.md,package.json' -e 'node_modules/*'")

		os.Exit(1)
	}

	// if !dryRun && len(filteredFiles) > 0 {
	// 	confirmAction(filteredFiles)
	// }

	output, err := createOutput(filteredFiles, prompt)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	fmt.Println("\nThe following files will be copied:")
	for _, file := range filteredFiles {
		fmt.Println(file)
	}

	if dryRun {
		fmt.Println("\nDry run, no action taken.")
		os.Exit(0)
	}

	if outputFile != "" {
		err = ioutil.WriteFile(outputFile, []byte(output), 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		}
		fmt.Println("Additionally, a file called '" + outputFile + "' was created.")
	}

	err = writeToClipboard(output)
	if err != nil {
		fmt.Println("Error writing to clipboard:", err)
		os.Exit(1)
	}

	fmt.Println("\n🥳 Copied to clipboard! Go paste it in ChatGPT.")
}
