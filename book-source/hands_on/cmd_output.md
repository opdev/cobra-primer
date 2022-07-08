# Writing to STDOUT and STDERR

It's fairly common to use `fmt.Println` to print output to the user. Both
`Println` and `Printf` will actually print to `os.Stdout` for us, and so it
serves as a convenience function for hacking on some code quickly, and getting
some text in front of the user.
[ref](https://cs.opensource.google/go/go/+/refs/tags/go1.18.3:src/fmt/print.go;l=273)

Cobra commands actually provide some wiring for printing things to the user via
Stdout or Stderr that allows us to configure its outputs to write to anything
that fulfills the io.Writer interface (e.g. log files, byte buffers, etc). We
don't need to concern ourselves with that quite yet, but what we do want to do
instead of using `fmt.Println` is use the built-in output for stdout/stderr if
we need them.

In our `sumCommandRunE`, we write the sum using `fmt.Println`:

```
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

	fmt.Println(sum(values...))         // Writes happen here!

	return nil
}
```

Instead, lets leverage the output target configured for the command. In
fairness, we haven't reconfigured it in this example, but we _can_, and that
will become more important when testing your cobra commands.

Change the line to look like this:

```go
	// fmt.Println(sum(values...))         // Old!
    fmt.Fprintln(cmd.OutOrStdout(), sum(values...))
```

If you're not familiar with Fprintln, it effectively allows you to _provide_ the
write target (the `io.Writer` interface) instead of assuming it should be
os.Stdout, as `fmt.Println` does. In this instance, we're passing in the cobra
command's configured writer. It, internally, will write to os.Stdout if nothing
else was configured.

We don't use it here, but there's an equivalent `cmd.OutOrStderr` function as
well.

If you run the command, nothing should have changed.

```text
$ go build -o math . && ./math sum 1 3
4
```