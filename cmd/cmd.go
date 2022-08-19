// Package cmd contains the terminal interface for the
// blockchain explorer.
package cmd

import (
	"fmt"
	"log"
	"strings"

	logparser "github.com/MalteHerrmann/BlockchainExplorer/parser"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

const (
	// Width of the height tile.
	blockHeightTileWidth = 18
)

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

	// Paragraph for block height
	HeightParagraph *widgets.Paragraph

	// Logparser for the blockchain information
	LP *logparser.LogParser
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
	myCmd.LP = lp

	// Create the basic grid that will hold all of the added
	// widgets.
	myCmd.View = ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	termWidthF := float64(termWidth)
	termHeightF := float64(termHeight)

	myCmd.URLParagraph = widgets.NewParagraph()
	myCmd.URLParagraph.Title = "Connection"
	myCmd.URLParagraph.Text = centerString(myCmd.URL, termWidth-2)
	myCmd.URLParagraph.TextStyle.Fg = ui.ColorGreen

	myCmd.HeightParagraph = widgets.NewParagraph()
	myCmd.HeightParagraph.Title = "Current Height"
	myCmd.HeightParagraph.Text = ""
	myCmd.HeightParagraph.TextStyle.Fg = ui.ColorYellow

	// Compose the layout
	URLRow := ui.NewRow(
		3.0/termHeightF,
		ui.NewCol(1.0, myCmd.URLParagraph),
	)
	BlockNumberRow := ui.NewRow(
		3.0/termHeightF,
		ui.NewCol(18.0/termWidthF, myCmd.HeightParagraph),
	)
	myCmd.View.Set(URLRow, BlockNumberRow)
	myCmd.View.SetRect(0, 0, termWidth, termHeight)

	return myCmd
}

// UpdateCmd updates the displayed information with the current
// information, that the logparser contains.
func (c *Cmd) UpdateCmd() {
	lastBlockNumber := c.LP.GetLastBlockNumber()
	c.HeightParagraph.Text = centerString(fmt.Sprint(lastBlockNumber), blockHeightTileWidth-2)
}

// centerString returns a string s that is centered given the total
// width w of the string.
func centerString(text string, width int) string {
	l := len(text)
	if l > width {
		return text[:width]
	}
	nFillersPerSide := (width - l) / 2
	return strings.Repeat(" ", nFillersPerSide) + text + strings.Repeat(" ", width-nFillersPerSide)
}
