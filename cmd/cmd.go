// Package cmd contains the terminal interface for the
// blockchain explorer.
package cmd

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

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
}

// NewCmd returns a new Cmd
func NewCmd(url string) *Cmd {
	myCmd := &Cmd{
		URL: url,
	}

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
