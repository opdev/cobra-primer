/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// arithmeticCmd represents the arithmetic command
var arithmeticCmd = &cobra.Command{
	Use:   "arithmetic",
	Short: "basic arithmetic functions",
}

func init() {
	rootCmd.AddCommand(arithmeticCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	arithmeticCmd.PersistentFlags().BoolP("show-inputs", "s", false, "whether to print inputs")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// arithmeticCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
