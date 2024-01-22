/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/hoehwa/meok/internal"
	"github.com/spf13/cobra"
)

// concurrencyCmd represents the concurrency command
var concurrencyCmd = &cobra.Command{
	Use:   "concurrency",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.PrettyPrint("/advanced/concurrency/")
	},
}

func init() {
	advancedCmd.AddCommand(concurrencyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// concurrencyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// concurrencyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
