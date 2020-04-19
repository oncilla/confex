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
	"encoding/json"
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	yml "gopkg.in/yaml.v3"

	"github.com/oncilla/confex/pkg/data"
)

// ConfigView renders the configuration.
type ConfigView struct {
	*ui.Grid
	tree     *widgets.Tree
	path     *widgets.Paragraph
	contents *widgets.Paragraph
	cfg      *data.Config
}

// NewConfigView constructs a new configuration view
func NewConfigView(cfg *data.Config) *ConfigView {
	tree := widgets.NewTree()
	tree.Title = " outline "
	tree.WrapText = false
	tree.SetNodes(newTree(cfg.Tree))
	tree.SelectedRowStyle = ui.NewStyle(ui.ColorGreen, ui.ColorClear, ui.ModifierBold)

	path := widgets.NewParagraph()
	path.Title = " path "

	contents := widgets.NewParagraph()
	contents.PaddingTop = 1
	contents.PaddingLeft = 1
	grid := ui.NewGrid()

	return &ConfigView{
		Grid:     grid,
		contents: contents,
		tree:     tree,
		cfg:      cfg,
		path:     path,
	}
}

func (cv *ConfigView) SetRect(x1, y1, x2, y2 int) {
	ratio := 3.0 / float64(y2-y1)
	cv.Grid = ui.NewGrid()

	cv.Grid.Set(
		ui.NewRow(1-ratio,
			ui.NewCol(0.25, cv.tree),
			ui.NewCol(0.75, cv.contents),
		),
		ui.NewRow(ratio,
			ui.NewCol(1, cv.path),
		),
	)

	cv.Grid.SetRect(x1, y1, x2, y2)

}

func (cv *ConfigView) Draw(b *ui.Buffer) {
	var contents interface{}
	var path string

	selected := cv.tree.SelectedNode()
	if selected != nil {
		contents = selected.Value.(*data.Node).Contents
		path = selected.Value.(*data.Node).Path()
	} else {
		contents = cv.cfg.Tree.Contents
	}
	raw, err := marshal(contents, cv.cfg.Language)
	if err != nil {
		// TODO(oncilla): add way to display error
		return
	}
	cv.contents.Text = string(raw) + " "
	cv.path.Text = path + " "

	if path != "" {
		cv.contents.Title = " " + path + " "
	} else {
		cv.contents.Title = " contents "
	}
	cv.Grid.Draw(b)

}

func marshal(contents interface{}, lang data.Language) ([]byte, error) {
	switch lang {
	case data.JSON:
		return json.MarshalIndent(contents, "", "  ")
	case data.YAML:
		return yml.Marshal(contents)
	case data.TOML:
		return encodeToml(contents)
	default:
		return nil, fmt.Errorf("unknown configuration language: %s", lang)
	}
}
