package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/vlazic/gpt-project-context/internal/clipboard"
	"github.com/vlazic/gpt-project-context/internal/fileutils"
)

func main() {
	include, exclude, prompt, outputFile, dryRun := getFlags()

	includePatterns := strings.Split(include, ",")
	excludePatterns := strings.Split(exclude, ",")

	filePaths, err := fileutils.GetFilePaths(includePatterns)
	if err != nil {
		fmt.Println("Error while walking the file tree:", err)
		return
	}

	filteredFiles := fileutils.FilterFiles(filePaths, excludePatterns)

	if len(filteredFiles) == 0 {
		fmt.Println("No files found.")
		// sample usage
		fmt.Println("Example usage Go: Usage: context -i '*.go,*.md' -e 'bin/*,*.txt'")
		fmt.Println("Example usage JS: Usage: context -i '*.js,README.md,package.json' -e 'node_modules/*'")

		os.Exit(1)
	}

	output, err := fileutils.CreateOutput(filteredFiles, prompt)
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
		err = fileutils.WriteToFile(outputFile, output)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		}
		fmt.Println("Additionally, a file called '" + outputFile + "' was created.")
	}

	err = clipboard.WriteToClipboard(output)
	if err != nil {
		fmt.Println("Error writing to clipboard:", err)
		os.Exit(1)
	}

	fmt.Println("\n🥳 Copied to clipboard! Go paste it in ChatGPT.")
}
