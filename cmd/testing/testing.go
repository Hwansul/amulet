/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package testing

import (
	"log"

	"github.com/hoehwa/meok/cmd"
	"github.com/spf13/cobra"
)

// TestingCmd represents the testing command
var TestingCmd = &cobra.Command{
	Use:   "testing",
	Short: "[Parent]A brief description of your command",
	Long: `Below are list of commands:
	
	golang: parent command to see subcommands for golang testing. 'meok testing golang'
	js: parent command to see subcommands for javascript testing 'meok testing js'
	`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(TestingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
