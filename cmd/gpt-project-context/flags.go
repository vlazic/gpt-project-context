package main

import "flag"

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
