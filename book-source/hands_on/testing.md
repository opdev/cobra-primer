# Testing

When testing cobra commands, you want to try and decouple your core logic into
libraries, similar to what we did with the `sum` and `subtract` functions for
their respective commands.

This makes testing a bit easier in that you don't need to wrap a bunch of
`cobra.Command` context into your unit tests.

With that said, you may find yourself needing to test calling your cobra command
to get more coverage. In that case, I would recommend borrowing a testing
function from the cobra library itself.

https://github.com/spf13/cobra/blob/main/command_test.go#L34-L43

```go
func executeCommandC(root *Command, args ...string) (c *Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}
```

This function calls your commands and returns your stdout/stderr streams
(together, but you can modify this to separate them if you need to), and an
error if your command returned one. It also returns your command should you need
it. 

So executing this for the subtract command would look something like this:

```go
func TestSubtractCmd(t *testing.T) {
	_, _, err := executeCommandC(rootCmd, "arithmetic", "subtract", "1", "2")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
```

Here, we pass `rootCmd` as our command to execute. We could also pass
`subtractCmd`, and then just pass `"1"` and `"2"` as parameters, but showing it
this way might help in understanding all the various ways you can leverage
`executeCommandC` to run your tests.

NOTE: Remember earlier that we wrote our execution output to the
`cmd.OutOrStdout` target. This is important here because we can actually
evaluate the output of our command execution. If we had used `fmt.Println`
instead, we would have a harder time trying to capture the command output stream
to evaluate for any failures.