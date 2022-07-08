package cmd

import (
	"testing"
	"bytes"
	"github.com/spf13/cobra"
)

func TestSubtractCmd(t *testing.T) {
	_, _, err := executeCommandC(rootCmd, "arithmetic", "subtract", "1", "2")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}