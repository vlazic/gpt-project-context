# Variables
BINARY_NAME ?= gpt-project-context
BINARY_DIR=bin
BINARY_PATH ?= $(BINARY_DIR)/$(BINARY_NAME)

.PHONY: build build-linux build-windows build-macos build-all release clean copy-context

build-linux:
	@echo "Building the Go binary..."
	@mkdir -p $(BINARY_DIR)
	@GOOS=linux GOARCH=amd64 go build -o $(BINARY_PATH)-linux main.go
	@echo "Build complete."

build-windows:
	@echo "Building the Go binary for Windows..."
	@mkdir -p $(BINARY_DIR)
	@GOOS=windows GOARCH=amd64 go build -o $(BINARY_DIR)/$(BINARY_NAME).exe main.go
	@echo "Build for Windows complete."

build-macos:
	@echo "Building the Go binary for macOS..."
	@mkdir -p $(BINARY_DIR)
	@GOOS=darwin GOARCH=amd64 go build -o $(BINARY_DIR)/$(BINARY_NAME)-macos main.go
	@echo "Build for macOS complete."

build-all: build build-windows build-macos

git-release: release
	npx standard-version --release-as $(new_version)
	git push --follow-tags origin master

release: build-all
	@echo "Creating a new GitHub release with the compiled binaries..."
	gh release create --title "GPT Project Context v$(new_version)" --notes-file CHANGELOG.md \
		--attach $(BINARY_DIR)/$(BINARY_NAME) --attach $(BINARY_DIR)/$(BINARY_NAME).exe \
		--attach $(BINARY_DIR)/$(BINARY_NAME)-macos
		--attach $(BINARY_DIR)/$(BINARY_NAME)-linux
	@echo "Release published."

clean:
	@echo "Cleaning up the build..."
	@rm -f $(BINARY_PATH)

copy-context: build
	$(BINARY_PATH) -i "*.go,Makefile,help.txt" -e "bin/*,context.txt"