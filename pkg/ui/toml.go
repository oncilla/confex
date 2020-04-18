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
	"bytes"
	"fmt"
	"reflect"

	"github.com/BurntSushi/toml"
)

// XXX(oncilla): This is super hacky.
func encodeToml(v interface{}) ([]byte, error) {
	k := reflect.ValueOf(v).Kind()
	switch k {
	case reflect.Map:
		buf := &bytes.Buffer{}
		err := toml.NewEncoder(buf).Encode(v)
		return buf.Bytes(), err
	case reflect.Slice:
		buf := &bytes.Buffer{}
		for _, sub := range v.([]interface{}) {
			e, err := encodeToml(sub)
			if err != nil {
				return nil, err
			}
			buf.Write(e)
			buf.WriteString("\n")
		}
		return buf.Bytes(), nil
	case reflect.String:
		return []byte(fmt.Sprintf(`"%s"`, v)), nil
	default:
		return []byte(fmt.Sprintf(`%v`, v)), nil
	}

}
