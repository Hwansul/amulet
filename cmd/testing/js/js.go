/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package js

import (
	"log"

	"github.com/hoehwa/meok/cmd/testing"
	"github.com/spf13/cobra"
)

// jsCmd represents the js command
var jsCmd = &cobra.Command{
	Use:   "js",
	Short: "A brief description of your command",
	Long: `Below are list of commands:
	
	anatomy: snippets to introduce how to test in javascript. 'meok testing js anatomy'
	backend: snippets and tips for backend testing. 'meok testing js backend'
	CI: short tips to test javascript program in CI Pipeline. 'meok testing js CI'
	coverage: explanations with snippets or images when to measure test coverage. 'meok testing js coverage'
	frontend: snippets and tips for backend testing. 'meok testing js frontend'
	tips: tips from 30-seconds-of-code. this is the only command come from other resources. 'meok testing js tips'
	`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	testing.TestingCmd.AddCommand(jsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
