package fileutils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestFilterFiles(t *testing.T) {
	// Setup
	tmpDir, err := setupMockProject()
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir) // Cleanup

	// .
	// ├── bin
	// │   └── binary_file
	// ├── src
	// │   ├── source.go
	// │   ├── source_test.go
	// │   └── helper.js
	// ├── main.go
	// ├── helper.go
	// ├── README.md
	// └── Makefile

	testCases := []struct {
		name              string
		includePatterns   []string
		excludePatterns   []string
		expectedFilePaths []string
	}{
		{
			name:              "include all .go and .md, Makefile, exclude src/source.go",
			includePatterns:   []string{"**/*.go", "**/*.md", "Makefile"},
			excludePatterns:   []string{"src/source.go"},
			expectedFilePaths: []string{"main.go", "helper.go", "README.md", "Makefile", "src/source_test.go"},
		},
		{
			name:              "include all .js files and Makefile",
			includePatterns:   []string{"**/*.js", "Makefile"},
			excludePatterns:   []string{},
			expectedFilePaths: []string{"src/helper.js", "Makefile"},
		},
		{
			name:              "include all .go files, exclude specific .go file",
			includePatterns:   []string{"**/*.go"},
			excludePatterns:   []string{"main.go"},
			expectedFilePaths: []string{"src/source.go", "helper.go", "src/source_test.go"},
		},
		{
			name:              "include all .go files, exclude .go files in src folder",
			includePatterns:   []string{"**/*.go"},
			excludePatterns:   []string{"src/*.go"},
			expectedFilePaths: []string{"main.go", "helper.go"},
		},
		{
			name:              "include all .go files, exclude *_test.go files",
			includePatterns:   []string{"**/*.go"},
			excludePatterns:   []string{"*/*_test.go"},
			expectedFilePaths: []string{"src/source.go", "main.go", "helper.go"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Generate a list of all file paths in the mock project that match the include patterns
			filePaths, err := GetFilePaths(tc.includePatterns, tmpDir)
			if err != nil {
				t.Fatal(err)
			}

			// Test FilterFiles
			filteredFiles := FilterFiles(filePaths, tc.excludePatterns)

			// Check if the filteredFiles slice contains the correct files
			if len(filteredFiles) != len(tc.expectedFilePaths) {
				t.Logf("expected files: %v", tc.expectedFilePaths)
				t.Logf("actual files: %v", filteredFiles)
				t.Fatalf("expected %d files, got %d", len(tc.expectedFilePaths), len(filteredFiles))
			}

			for _, expectedPath := range tc.expectedFilePaths {
				found := false
				for _, filteredFile := range filteredFiles {
					if expectedPath == filteredFile {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("expected file not found: %s", expectedPath)
				}
			}
		})
	}
}

func setupMockProject() (string, error) {
	// Create a temporary directory
	tmpDir, err := ioutil.TempDir("", "test_project")
	if err != nil {
		return "", err
	}

	// Define the file structure:
	// 	.
	// ├── bin
	// │   └── binary_file
	// ├── src
	// │   ├── source.go
	// │   ├── source_test.go
	// │   └── helper.js
	// ├── main.go
	// ├── helper.go
	// ├── README.md
	// └── Makefile
	files := map[string]string{
		"bin/binary_file":    "",
		"src/source.go":      "package src",
		"src/source_test.go": "package main",
		"src/helper.js":      "console.log('Hello, world!');",
		"main.go":            "package main",
		"helper.go":          "package main",
		"README.md":          "# Test Project",
		"Makefile":           "",
	}

	// Create the files
	for path, content := range files {
		fullPath := filepath.Join(tmpDir, path)
		err := os.MkdirAll(filepath.Dir(fullPath), 0755)
		if err != nil {
			return "", err
		}
		err = ioutil.WriteFile(fullPath, []byte(content), 0644)
		if err != nil {
			return "", err
		}
	}

	return tmpDir, nil
}
