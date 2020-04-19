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

package data

import (
	"fmt"
	"sort"
	"strings"
)

// Node describes a node in the configuration
type Node struct {
	Name     string
	Contents interface{}
	Parent   *Node
	Nodes    []*Node
}

// NewTree constructs a new tree from the parsed config.
func NewTree(y interface{}) (*Node, error) {
	node := &Node{
		// Is set by parent, if any.
		Name:     "root",
		Contents: y,
	}
	switch v := y.(type) {
	case bool, string, int, uint, int64, float64, nil:
		return node, nil
	case map[string]interface{}:
		for key, child := range v {
			cnode, err := NewTree(child)
			if err != nil {
				return nil, fmt.Errorf("at %s: %w", key, err)
			}
			cnode.Name = key
			cnode.Parent = node
			node.Nodes = append(node.Nodes, cnode)
		}
		sort.Slice(node.Nodes, func(i, j int) bool {
			return node.Nodes[i].Name < node.Nodes[j].Name
		})
		return node, nil
	case []interface{}:
		for i, child := range v {
			cnode, err := NewTree(child)
			if err != nil {
				return nil, fmt.Errorf("at %d: %w", i, err)
			}
			cnode.Name = fmt.Sprintf("[%d]", i)
			cnode.Parent = node
			node.Nodes = append(node.Nodes, cnode)
		}
		return node, nil
	default:
		return nil, fmt.Errorf("unhandled type %T", v)
	}
}

// Path specifies the path to this node.
func (n *Node) Path() string {
	var parts []string
	for n.Parent != nil {
		parts = append([]string{n.Name}, parts...)
		n = n.Parent
	}
	return strings.Join(parts, ".")
}

func (n *Node) String() string {
	return n.Name + " "
}
