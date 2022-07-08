/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// supersecretmathCmd represents the supersecretmath command
var supersecretmathCmd = &cobra.Command{
	Use:    "supersecretmath",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is where we do super secret math!")
	},
}

func init() {
	rootCmd.AddCommand(supersecretmathCmd)
}
