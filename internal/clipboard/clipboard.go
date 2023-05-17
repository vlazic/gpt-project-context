package clipboard

import (
	"context"
	"fmt"
	"time"

	"golang.design/x/clipboard"
)

func WriteToClipboard(output string) error {
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
