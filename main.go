// Package TerminalExplorer contains the terminal user interface for the
// blockchain explorer. It is an alternative to the graphical
// user interface.
//
// The script is expected to be called with the URL of the server to
// connect to as the first argument.
package main

import (
	"log"
	"os"
	"time"

	"github.com/MalteHerrmann/terminalexplorer/cmd"
	ui "github.com/gizak/termui/v3"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	url := os.Args[1]
	terminalExplorer := cmd.NewCmd(url)
	ui.Render(terminalExplorer.View)

	go func() {
		terminalExplorer.LP.SubscribeToBlocks()
	}()

	go func() {
		for range time.Tick(time.Second) {
			terminalExplorer.UpdateCmd()
			ui.Render(terminalExplorer.View)
		}
	}()

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
