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
	tree.Title = "tree"
	tree.WrapText = false
	tree.SetNodes(newTree(cfg.Tree))
	tree.SelectedRowStyle = ui.NewStyle(ui.ColorGreen, ui.ColorClear, ui.ModifierBold)

	path := widgets.NewParagraph()
	path.Title = "path"

	contents := widgets.NewParagraph()
	contents.Title = "contents"

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		ui.NewRow(0.87,
			ui.NewCol(0.25, tree),
			ui.NewCol(0.75, contents),
		),
		ui.NewRow(0.13,
			ui.NewCol(1, path),
		),
	)
	return &ConfigView{
		Grid:     grid,
		contents: contents,
		tree:     tree,
		cfg:      cfg,
		path:     path,
	}
}

func (cv *ConfigView) render() error {
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
		return err
	}
	cv.contents.Text = string(raw) + " "
	cv.path.Text = path
	ui.Render(cv)
	return nil
}

func marshal(contents interface{}, lang data.Language) ([]byte, error) {
	switch lang {
	case data.JSON:
		return json.MarshalIndent(contents, "", "    ")
	case data.YAML:
		return yml.Marshal(contents)
	case data.TOML:
		return encodeToml(contents)
	default:
		return nil, fmt.Errorf("unknown configuration language: %s", lang)
	}
}
