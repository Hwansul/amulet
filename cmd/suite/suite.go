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

package suite

import (
	"fmt"

	"github.com/jipilmuk/amulet/cmd"
	"github.com/spf13/cobra"
)

// suiteCmd represents the suite command.
var suiteCmd = &cobra.Command{
	Use:   "suite",
	Short: "[Parent]Dev-suite commands",
	Long: `Below are list of commands:

	commandLine: snippets to develop command line program 'amulet suite commandLine'
	file: snippets to treat file in program 'amulet suite file'
	http: snippets to treat http request and responses 'amulet suite http'
	path: snippets to treat path in program  'amulet suite path'
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
		Run 'help suite <command>' for more information on a specific command.
		
		you can use following sub commands:
		- amulet suite path
		- amulet suite http
		- amulet suite file
		- amulet suite commandLine
		`)
	},
}

func init() {
	cmd.RootCmd.AddCommand(suiteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// suiteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// suiteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
