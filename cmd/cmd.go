// Package cmd contains the terminal interface for the
// blockchain explorer.
package cmd

import (
	"fmt"
	"log"

	logparser "github.com/MalteHerrmann/BlockchainExplorer/parser"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// TODO: Add logparser to update infos
// TODO: Add more infos to be displayed
// TODO: Add tests

// Cmd is the type for the terminal interface
type Cmd struct {
	// URL of the server to connect to
	URL string

	// View is the grid containing all widgets
	View *ui.Grid

	// Paragraph for connected URL
	URLParagraph *widgets.Paragraph

	// Paragraph for block number
	BlockParagraph *widgets.Paragraph

	// Logparser for the blockchain information
	lp *logparser.LogParser
}

// NewCmd creates a new Cmd and populates it with the default
// widgets. After setting up the view, it returns the Cmd.
func NewCmd(url string) *Cmd {
	myCmd := &Cmd{
		URL: url,
	}

	lp, err := logparser.NewLogParserWithURL(myCmd.URL)
	if err != nil {
		log.Fatalf("Error creating logparser: %v", err)
	}
	myCmd.lp = lp

	// Create the basic grid that will hold all of the added
	// widgets.
	myCmd.View = ui.NewGrid()

	myCmd.URLParagraph = widgets.NewParagraph()
	myCmd.URLParagraph.Text = "Connected to " + myCmd.URL

	myCmd.BlockParagraph = widgets.NewParagraph()
	myCmd.BlockParagraph.Text = "Block number"

	URLRow := ui.NewRow(1.0/2.0, ui.NewCol(1.0, myCmd.URLParagraph))
	BlockNumberRow := ui.NewRow(1.0/2.0, ui.NewCol(1.0, myCmd.BlockParagraph))
	myCmd.View.Set(URLRow, BlockNumberRow)
	termWidth, termHeight := ui.TerminalDimensions()
	myCmd.View.SetRect(0, 0, termWidth, termHeight)

	return myCmd
}

// UpdateCmd updates the displayed information with the current
// information, that the logparser contains.
func UpdateCmd(*Cmd) {
	fmt.Println("Updating.")
}
