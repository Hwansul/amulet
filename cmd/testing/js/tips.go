/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package js

import (
	"github.com/hoehwa/meok/utills"
	"github.com/spf13/cobra"
)

// tipsCmd represents the tips command
var tipsCmd = &cobra.Command{
	Use:   "tips",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		utills.PrintContent("/testing/js/tips")
	},
}

func init() {
	jsCmd.AddCommand(tipsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tipsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tipsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
