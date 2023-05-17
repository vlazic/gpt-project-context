# Variables
BINARY_NAME ?= gpt-project-context
BINARY_DIR=bin
BINARY_PATH ?= $(BINARY_DIR)/$(BINARY_NAME)

.PHONY: build build-linux build-windows build-macos build-all release clean copy-context

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

github-release: build-all
	sed -r -i "s/(vlazic\/$(BINARY_NAME)\/releases\/download\/)[^\/]+\//\1v$(new_version)\//g" README.md

	npx standard-version --release-as $(new_version)
	git push --follow-tags origin master

	@echo "Creating a new GitHub release with the compiled binaries..."
	gh release create "v$(new_version)" -F CHANGELOG.md $(BINARY_DIR)/*

	@echo "Release published."

clean:
	@echo "Cleaning up the build..."
	@rm -f $(BINARY_PATH)

copy-context: build-linux
	$(BINARY_PATH)-linux -i "*.go,Makefile,README.md,go.mod" -e "bin/*,context.txt"
