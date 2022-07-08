/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// subtractCmd represents the subtract command
var subtractCmd = &cobra.Command{
	Use:   "subtract",
	Short: "subtract integers",
	Args:  cobra.ExactArgs(2),
	RunE:  subtractCommandRunE,
}

func subtractCommandRunE(cmd *cobra.Command, args []string) error { // return an error
	// convert args which is []string to []int
	values := make([]int, len(args))
	for i, v := range args {
		vAsInt, err := strconv.Atoi(v)
		if err != nil { // new code!
			return fmt.Errorf("you provide a value that was not an integer: %s", v)
		}
		values[i] = vAsInt
	}

	result := subtract(values...)
	invert, _ := cmd.Flags().GetBool("invert-sign")
	if invert {
		result = -result
	}

	fmt.Fprintln(cmd.OutOrStdout(), result)

	return nil
}

func subtract(values ...int) int {
	x := values[0]
	for _, v := range values[1:] {
		x -= v
	}

	return x
}

func init() {
	arithmeticCmd.AddCommand(subtractCmd)
	subtractCmd.Flags().BoolP("invert-sign", "i", false, "inverts the sign of the result.")
}
