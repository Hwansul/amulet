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

package cmd

import (
	"log"
	"os"
	"strings"

	listFancy "github.com/hwansul/amulet/lib/bubbletea/list-fancy"
	constant "github.com/hwansul/amulet/lib/constant"
	"github.com/spf13/cobra"
)

// stringCmd represents the string command.
var stringCmd = &cobra.Command{
	Use:   "string",
	Short: "Print list of snippets for string.",
	Long: `Fuzzy find and print list of snippets for string.
	
supported languages:
python
javascript
golang
bash
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var titles []string
		var descs []string

		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		stringDir := userHomeDir + "/snippet/string"
		files, err := os.ReadDir(stringDir)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			titles = append(titles, file.Name())
			descs = append(descs, constant.GardenBaseurl+"/flower/snippet/string"+strings.Split(file.Name(), ".")[0])
		}

		listFancy.Init(titles, descs, stringDir)
	},
}

func init() {
	rootCmd.AddCommand(stringCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stringCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stringCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
