/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package golang

import (
	"log"

	"github.com/hoehwa/meok/cmd/testing"
	"github.com/spf13/cobra"
)

// golangCmd represents the golang command
var golangCmd = &cobra.Command{
	Use:   "golang",
	Short: "A brief description of your command",
	Long: `Below are list of commands:
	
	acptnTest: documents and snippets for acceptance testing. 'meok testing golang acptnTest'
	antiPtrns: documents and snippets for anti patterns during testing. 'meok testing golang antiPtrns'
	scaleAcptn: documents and snippets that describes how to scale up acceptance in golang testing. 'meok testing golang scaleAcptn'
	tips: documents and snippets for tips. 'meok testing golang tips'
	withoutMocks: documents and snippets which explain why needs to do testing without. 'meok testing golang withoutMocks'
	`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	testing.TestingCmd.AddCommand(golangCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// golangCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// golangCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
