package cmd

import (
	"github.com/hoehwa/meok/utills"
	"github.com/spf13/cobra"
)

// dataprocessCmd represents the dataprocess command.
var dataprocessCmd = &cobra.Command{
	Use:   "dataprocess",
	Short: "[root] snippets for dataprocess",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		utills.PrintContent("/dataprocess")
	},
}

func init() {
	RootCmd.AddCommand(dataprocessCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dataprocessCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dataprocessCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
