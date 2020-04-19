// MIT License
//
// Copyright (c) 2020 Oncilla
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package ui

import (
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/oncilla/confex/pkg/data"
)

const (
	displayContents = iota
	displayHelp
)

// Window is the main container for displaying.
type Window struct {
	help struct {
		tabs *widgets.TabPane
		text *widgets.Paragraph
	}
	explorer struct {
		tabs     *widgets.TabPane
		contents *ConfigView
	}
	displayMode int
	exit        bool
}

// NewWindow constructs a new window.
func NewWindow(cfg *data.Config) *Window {
	helpTab := widgets.NewTabPane("<ANY> Press to return")
	helpTab.ActiveTabStyle = ui.StyleClear
	helpTab.Border = false

	helpText := widgets.NewParagraph()
	helpText.Text = helpMsg
	helpText.PaddingLeft = 1

	explorerTab := widgets.NewTabPane("<ENTER>,e: Toggle expand", "E: Expand all", "C: Collapse all", "h: Help", "q: quit")
	explorerTab.ActiveTabStyle = ui.StyleClear
	explorerTab.Border = false

	w := &Window{}
	w.help.tabs = helpTab
	w.help.text = helpText
	w.explorer.tabs = explorerTab
	w.explorer.contents = NewConfigView(cfg)
	return w
}

// Run starts the process.
func (w *Window) Run() error {
	if err := ui.Init(); err != nil {
		return err
	}
	defer ui.Close()

	termWidth, termHeight := ui.TerminalDimensions()
	w.SetRect(0, 0, termWidth, termHeight)
	w.Render()

	previousKey := ""
	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(500 * time.Millisecond)

	for !w.exit {
		select {
		case e := <-uiEvents:

			switch e.ID {
			case "g":
				if previousKey == "g" {
					e.ID = "gg"
					w.handleEvent(e)
				} else {
					w.handleEvent(e)
				}
			case "<Resize>":
				width, height := ui.TerminalDimensions()
				w.SetRect(0, 0, width, height)
				ui.Clear()
			default:
				w.handleEvent(e)
			}

			// Clear key state.
			if previousKey == "g" {
				previousKey = ""
			} else {
				previousKey = e.ID
			}
		case <-ticker.C:
		}

		w.Render()
	}
	return nil
}

func (w *Window) handleEvent(e ui.Event) {
	switch w.displayMode {
	case displayHelp:
		w.displayMode = displayContents
	default:
		w.handleEventExplorer(e)
	}
}

func (w *Window) handleEventExplorer(e ui.Event) {
	switch e.ID {
	case "q", "<C-c>", "Q":
		w.exit = true
	case "j", "<Down>":
		w.explorer.contents.tree.ScrollDown()
	case "k", "<Up>":
		w.explorer.contents.tree.ScrollUp()
	case "<C-d>":
		w.explorer.contents.tree.ScrollHalfPageDown()
	case "<C-u>":
		w.explorer.contents.tree.ScrollHalfPageUp()
	case "<C-f>":
		w.explorer.contents.tree.ScrollPageDown()
	case "<C-b>":
		w.explorer.contents.tree.ScrollPageUp()
	case "gg":
		w.explorer.contents.tree.ScrollTop()
	case "<Home>":
		w.explorer.contents.tree.ScrollTop()
	case "e", "<Enter>":
		w.explorer.contents.tree.ToggleExpand()
	case "G", "<End>":
		w.explorer.contents.tree.ScrollBottom()
	case "E":
		w.explorer.contents.tree.ExpandAll()
	case "c", "C":
		// TODO(oncilla): Figure out why this is needed.
		w.explorer.contents.tree.SelectedRow = 0
		w.explorer.contents.tree.CollapseAll()
	case "h", "H":
		w.displayMode = displayHelp
	}
}

// SetRect rearranges the window.
func (w *Window) SetRect(x1, y1, x2, y2 int) {
	w.help.text.SetRect(x1, y1, x2, y2-1)
	w.help.tabs.SetRect(0, y2-1, x2, y2)
	w.explorer.contents.SetRect(x1, y1, x2, y2-1)
	w.explorer.tabs.SetRect(0, y2-1, x2, y2)
}

// Render renders the window.
func (w *Window) Render() {
	switch w.displayMode {
	case displayContents:
		ui.Render(w.explorer.tabs, w.explorer.contents)
	case displayHelp:
		ui.Render(w.help.tabs, w.help.text)
	}
}
