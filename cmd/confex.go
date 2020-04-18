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

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/oncilla/confex/pkg/data"
	"github.com/oncilla/confex/pkg/ui"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var rootFlags struct {
}

// Root is the entry point for the confex application.
var Root = &cobra.Command{
	Use:   "confex [file]",
	Short: "A terminal based configuration file explorer",
	Args:  cobra.MaximumNArgs(1),
	// See https://github.com/spf13/cobra/issues/340#issuecomment-374617413.
	Long: `A terminal based configuration file explorer

You can pass in any json, yaml or toml file and explore it interactively.
`,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO(oncilla): support piped files.
		cmd.SilenceUsage = true

		cfg, err := fromFile(args[0])
		if err != nil {
			return err
		}
		return ui.ControlLoop(cfg)
	},
}

func fromFile(name string) (*data.Config, error) {
	raw, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return tryAll(raw)
}

func tryAll(raw []byte) (*data.Config, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(raw, &m); err == nil {
		tree, err := data.NewTree(m)
		if err != nil {
			return nil, err
		}
		return &data.Config{
			Tree:     tree,
			Language: data.JSON,
		}, nil
	}
	if err := yaml.Unmarshal(raw, &m); err == nil {
		tree, err := data.NewTree(m)
		if err != nil {
			return nil, err
		}
		return &data.Config{
			Tree:     tree,
			Language: data.YAML,
		}, nil
	}
	if err := toml.Unmarshal(raw, &m); err == nil {
		tree, err := data.NewTree(m)
		if err != nil {
			return nil, err
		}
		return &data.Config{
			Tree:     tree,
			Language: data.TOML,
		}, nil
	}
	return nil, fmt.Errorf("could not identify format")
}
