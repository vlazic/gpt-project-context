# gpt-project-context

`gpt-project-context` is a command-line tool designed to help developers quickly provide context about their project as initial input to ChatGPT. It scans the project files and generates a text output that includes code snippets, file structures, and other relevant information, which can be easily shared with AI language models like OpenAI's ChatGPT.

## Contributing

We welcome any issues and pull requests in the spirit of having mostly GPT build out this tool. Using [ChatGPT Plus](https://chat.openai.com/) is recommended for quick access to GPT-4 and getting the best results.

## Getting Started

To get started with `gpt-project-context`, follow these steps:

1. Ensure you have Go installed on your system.
2. Clone or download the `gpt-project-context` repository.
3. Navigate to the repository's root directory in your terminal.
4. Run `make build` to build the Go binary.
5. Use the generated `gpt-project-context` binary with the following command:

   ```bash
   ./bin/gpt-project-context -i "*.go,*.md" -e "bin/*,*.txt"
   ```

   Replace the include and exclude patterns as needed for your project. The `-i` flag specifies the file types to include, and the `-e` flag specifies the patterns to exclude.

6. The tool will generate a text output containing the project context. This output will be automatically copied to your clipboard, ready for use with AI language models like ChatGPT.

## Usage Examples

Here are some usage examples for different programming languages:

- Go: `./bin/gpt-project-context -i "*.go,*.md" -e "bin/*,*.txt"`
- JavaScript: `./bin/gpt-project-context -i "*.js,README.md,package.json" -e "node_modules/*"`

Feel free to adjust the include and exclude patterns to match the specific needs of your project.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
