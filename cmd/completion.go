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
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var completionShell string

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates bash completion scripts",
	Long: `'completion' outputs the autocomplete configuration for some shells.

For example, you can add autocompletion for your current bash session using:

    . <( confex completion )

To permanently add bash autocompletion, run:

    confex completion > /etc/bash_completion.d/confex
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		switch completionShell {
		case "bash":
			return Root.GenBashCompletion(os.Stdout)
		case "zsh":
			return Root.GenZshCompletion(os.Stdout)
		case "fish":
			return Root.GenFishCompletion(os.Stdout, true)
		default:
			return fmt.Errorf("unknown shell: %s", completionShell)
		}
	},
}

func init() {
	Root.AddCommand(completionCmd)
	completionCmd.Flags().StringVar(&completionShell, "shell", "bash",
		"Shell type (bash|zsh|fish)")
}
