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

package syntax

import (
	"fmt"

	"github.com/hwansul/amulet/cmd"
	"github.com/spf13/cobra"
)

// syntaxCmd represents the syntax command.
var syntaxCmd = &cobra.Command{
	Use:   "syntax",
	Short: "Snippets for programming syntax.",
	Long: `Below are list of commands:
	
	errorHandling: snippets for error handling syntax 'amulet syntax errorHandling'
	function: snippets for function syntax 'amulet syntax function'
	iterable: snippets for iterables. For example, array in js, list in python, etc... 'amulet syntax iterable'
	keyValue: snippets for key-value pairs syntax. For example, object in js, dictionary in python, etc...  'amulet syntax keyValue'
	type: snippets for data type. For example, union in ts, type hint in python, etc...  'amulet syntax type'
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
		Run 'help suite <command>' for more information on a specific command.
		
		you can use following sub commands:
		- amulet syntax errorHandling
		- amulet syntax function
		- amulet syntax iterable
		- amulet syntax keyValue
		- amulet syntax type
		`)
	},
}

func init() {
	cmd.RootCmd.AddCommand(syntaxCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syntaxCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syntaxCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
