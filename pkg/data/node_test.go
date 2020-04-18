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

package data_test

import (
	"testing"

	"github.com/oncilla/confex/pkg/data"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTree(t *testing.T) {
	input := map[string]interface{}{
		"a": map[string]interface{}{
			"c": "hello",
			"d": []interface{}{
				"preserve",
				1337,
				"order",
			},
			"b": 1,
		},
	}

	tree, err := data.NewTree(input)
	require.NoError(t, err)

	assert.Equal(t, "root", tree.Name)
	assert.Len(t, tree.Nodes, 1)
	assert.Equal(t, input, tree.Contents)
	assert.Nil(t, tree.Parent)
	assert.Empty(t, tree.Path())

	a := tree.Nodes[0]
	assert.Equal(t, "a", a.Name)
	assert.Len(t, a.Nodes, 3)
	assert.Equal(t, input["a"], a.Contents)
	assert.Equal(t, tree, a.Parent)
	assert.Equal(t, "a", a.Path())

	b := a.Nodes[0]
	assert.Equal(t, "b", b.Name)
	assert.Len(t, b.Nodes, 0)
	assert.Equal(t, 1, b.Contents)
	assert.Equal(t, a, b.Parent)
	assert.Equal(t, "a.b", b.Path())

	c := a.Nodes[1]
	assert.Equal(t, "c", c.Name)
	assert.Len(t, c.Nodes, 0)
	assert.Equal(t, "hello", c.Contents)
	assert.Equal(t, a, c.Parent)
	assert.Equal(t, "a.c", c.Path())

	d := a.Nodes[2]
	assert.Equal(t, "preserve", d.Nodes[0].Contents)
	assert.Equal(t, 1337, d.Nodes[1].Contents)
	assert.Equal(t, "order", d.Nodes[2].Contents)
}
