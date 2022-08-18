// Package TerminalExplorer contains the terminal user interface for the
// blockchain explorer. It is an alternative to the graphical
// user interface.
package main

import (
	"log"

	"github.com/MalteHerrmann/terminalexplorer/cmd"
	ui "github.com/gizak/termui/v3"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	url := "http://localhost:8000"

	terminalExplorer := cmd.NewCmd(url)
	ui.Render(terminalExplorer.View)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
