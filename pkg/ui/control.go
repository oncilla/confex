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
	"fmt"
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/oncilla/confex/pkg/data"
)

// ControlLoop controls the terminal UI.
func ControlLoop(cfg *data.Config) error {
	if err := ui.Init(); err != nil {
		return err
	}
	defer ui.Close()

	cv := NewConfigView(cfg)
	_ = cv.render()

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return nil
		case "j", "<Down>":
			cv.tree.ScrollDown()
		case "k", "<Up>":
			cv.tree.ScrollUp()
		case "<C-d>":
			cv.tree.ScrollHalfPageDown()
		case "<C-u>":
			cv.tree.ScrollHalfPageUp()
		case "<C-f>":
			cv.tree.ScrollPageDown()
		case "<C-b>":
			cv.tree.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				cv.tree.ScrollTop()
			}
		case "<Home>":
			cv.tree.ScrollTop()
		case "<Enter>":
			cv.tree.ToggleExpand()
		case "G", "<End>":
			cv.tree.ScrollBottom()
		case "E":
			cv.tree.ExpandAll()
		case "C":
			// TODO(oncilla): Figure out why this is needed.
			cv.tree.SelectedRow = 0
			cv.tree.CollapseAll()
		case "<Resize>":
			payload := e.Payload.(ui.Resize)
			cv.SetRect(0, 0, payload.Width, payload.Height)
			ui.Clear()
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		// FIXME(oncilla): This should write to an error box instead.
		if err := cv.render(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}
}
