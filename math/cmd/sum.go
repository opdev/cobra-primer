/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// sumCmd represents the sum command
var sumCmd = &cobra.Command{
	Use:        "sum",
	Aliases:    []string{"total"},
	Deprecated: `This command will be replaced by the "addition" command in the next release`,
	Args:       cobra.MinimumNArgs(1),
	Short:      "add integers",
	Long: `Given an arbitrary number of integer arguments,
	
this will return the sum of all values.`,

	RunE: sumCommandRunE,
}

func sumCommandRunE(cmd *cobra.Command, args []string) error { // return an error
	// convert args which is []string to []int
	values := make([]int, len(args))
	for i, v := range args {
		vAsInt, err := strconv.Atoi(v)
		if err != nil { // new code!
			return fmt.Errorf("you provide a value that was not an integer: %s", v)
		}
		values[i] = vAsInt
	}

	showInputs, _ := cmd.Flags().GetBool("show-inputs")
	if showInputs {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n", strings.Join(args, "+"))
	}

	fmt.Fprintln(cmd.OutOrStdout(), sum(values...))

	return nil
}

func sum(values ...int) int {
	x := 0
	for _, v := range values {
		x += v
	}

	return x
}

func init() {
	arithmeticCmd.AddCommand(sumCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
