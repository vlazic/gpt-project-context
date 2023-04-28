# GPT Project Context

<!-- ![GitHub release (latest by date)](https://img.shields.io/github/v/release/vlazic/gpt-project-context)] -->

`gpt-project-context` is a command-line tool designed to boost developers' productivity by facilitating seamless interaction with AI language models like OpenAI's ChatGPT. By scanning project files and generating a comprehensive text output consisting of code snippets, file structures, and other relevant details, developers can easily share crucial context about their projects with AI language models.

I, as the author, use this tool for every project, and I'm confident it will help you speed up your work with projects too.

## Installation

To quickly install `gpt-project-context` using binaries from GitHub release, follow the instructions for your operating system:

### macOS

```sh
# Download the binary for macOS
curl -L -o gpt-project-context "https://github.com/vlazic/gpt-project-context/releases/download/v1.0.1/gpt-project-context-macos"

# Make it executable
chmod +x gpt-project-context

# Move it to your PATH
sudo mv gpt-project-context /usr/local/bin/
```

### Windows

1. Download the `.exe` file from the [GitHub releases page](https://github.com/vlazic/gpt-project-context/releases).
2. Move the `.exe` file to a folder included in your `PATH` environment variable (e.g., `C:\Windows\System32`).

### Linux

```sh
# Download the binary for Linux
curl -L -o gpt-project-context "https://github.com/vlazic/gpt-project-context/releases/download/v1.0.1/gpt-project-context-linux"

# Make it executable
chmod +x gpt-project-context

# Move it to your PATH
sudo mv gpt-project-context /usr/local/bin/
```

## Usage

### Include and Exclude Flags

`gpt-project-context` allows you to customize the files included in the output using the `-i` (include) and `-e` (exclude) flags. The `-i` flag specifies which file patterns to include, while the `-e` flag specifies which file patterns to exclude.

All available flags are listed below:

```sh
Usage of ./bin/gpt-project-context:
  -e string
        exclude patterns
  -i string
        include patterns
  -n    no action, do not copy or write to clipboard
  -o string
        output file path
  -p string
        prompt at the beginning (default "Here is the context of my current project. Just respond with 'OK' and wait for the instructions:")
```

### Examples

#### Go:

```sh
gpt-project-context -i '*.go,*.md' -e 'bin/*,*.txt'
```

#### JavaScript:

```sh
gpt-project-context -i '*.js,README.md,package.json' -e 'node_modules/*'
```

To use this tool more conveniently in a JavaScript project, add it as an npm run script in your `package.json`:

```json
{
  "scripts": {
    "context": "gpt-project-context -i '*.js,README.md,package.json' -e 'node_modules/*'"
  }
}
```

Now, you can simply run `npm run context` to execute the script.

## Contributing

We welcome any issues and pull requests. If you have any questions, please feel free to open an issue.

## License

This project is licensed under the [MIT License](LICENSE).
