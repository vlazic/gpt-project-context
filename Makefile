# Variables
BINARY_NAME ?= gpt-project-context
BINARY_DIR=bin
BINARY_PATH ?= $(BINARY_DIR)/$(BINARY_NAME)

.PHONY: build-linux build-windows build-macos build-all clean copy-context github-release help test

help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  build-linux    to build the Go binary for Linux"
	@echo "  build-windows  to build the Go binary for Windows"
	@echo "  build-macos    to build the Go binary for macOS"
	@echo "  build-all      to build the Go binary for all platforms"
	@echo "  github-release to create a new GitHub release with the compiled binaries. Usage: make github-release new_version=<version>"
	@echo "  copy-context   to create a context of this project for the GPT"
	@echo "  clean          to remove the built Go binaries"
	@echo "  help           to display this help message"
	@echo "  test           to run the tests"

clean:
	@echo "Cleaning up the build..."
	@rm -f $(BINARY_PATH)

test:
	go test -v ./...

build-linux:
	@echo "Building the Go binary..."
	@mkdir -p $(BINARY_DIR)
	@GOOS=linux GOARCH=amd64 go build -o $(BINARY_PATH)-linux ./cmd/gpt-project-context
	@echo "Build complete."

build-windows:
	@echo "Building the Go binary for Windows..."
	@mkdir -p $(BINARY_DIR)
	@GOOS=windows GOARCH=amd64 go build -o $(BINARY_DIR)/$(BINARY_NAME).exe ./cmd/gpt-project-context
	@echo "Build for Windows complete."

build-macos:
	@echo "Building the Go binary for macOS..."
	@mkdir -p $(BINARY_DIR)
	@GOOS=darwin GOARCH=amd64 go build -o $(BINARY_DIR)/$(BINARY_NAME)-macos ./cmd/gpt-project-context
	@echo "Build for macOS complete."

build-all: build-linux build-windows build-macos

github-release: clean test build-all
	sed -r -i "s/(vlazic\/$(BINARY_NAME)\/releases\/download\/)[^\/]+\//\1v$(new_version)\//g" README.md

	npx standard-version --release-as $(new_version)
	git add README.md
	git commit --amend --no-edit
	# git push --follow-tags origin master

	@echo "Creating a new GitHub release with the compiled binaries..."
	gh release create "v$(new_version)" -F CHANGELOG.md $(BINARY_DIR)/*

	@echo "Release published."


copy-context: build-linux
	$(BINARY_PATH)-linux -i "*.go,Makefile,README.md,go.mod" -e "bin/*,context.txt"
