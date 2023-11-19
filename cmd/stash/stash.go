// Copyright © 2023 Mindulle <mindullestudio@gmail.com>

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package stash

import (
	"fmt"

	"github.com/hwansul/amulet/cmd"
	"github.com/spf13/cobra"
)

// stashCmd represents the stash command.
var stashCmd = &cobra.Command{
	Use:   "stash",
	Short: "[Subcommands]To treat special things or concepts.",
	Long: `Below are list of commands:
	
	async: snippets for asynchronous communication syntax 'amulet stash async'
	browser: snippets to treat browser behaviors 'amulet stash browser'
	concurrency: snippets for concurrency control. For example, goroutine in golang, etc... 'amulet stash concurrency'
	testing: snippets for testing. 'amulet stash testing'
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
		Run 'help suite <command>' for more information on a specific command.
		
		you can use following sub commands:
		- amulet stash async
		- amulet stash browser
		- amulet stash concurrency
		- amulet stash testing
		`)
	},
}

func init() {
	cmd.RootCmd.AddCommand(stashCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stashCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stashCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
