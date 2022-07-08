# Adding a flag

So far, we've been working with subcommands and positional args. For example:

```
COMMAND SUBCOMMAND POSITIONALARGS...
math    subtract   1 2
```

Cobra documentation suggests that your **subcommands** should describe your
actions, and **flags** should modify those actions. So for the subtract command,
we're going to add a flag that inverts the sign of the integer. Lets call it
`--invert-sign`. Our result would look like:

```
math subtract --invert-sign 1 2
1
```

It's a pretty silly example, but we don't want to spend time on our logic. We
want to spend time on cobra!

So let's implement a flag in the `subtract.go` file. Take a look at the `init`
section at the very bottom. It probably contains some commented code:

```go
func init() {
	rootCmd.AddCommand(subtractCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subtractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subtractCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
```

It also already contains our `AddCommand` call, binding the `subtractCmd` to our
`rootCmd`. We didn't have to do any of that wiring - it was done for us by
`cobra-cli`.

We're going to create a `local` flag, and so we're going to reuse the last line
of this commented code. Go ahead and uncomment this line. Replace `toggle` with
`invert-sign`, and `t` with `i`. 

```go
func init() {
	rootCmd.AddCommand(subtractCmd)
    // ...
	subtractCmd.Flags().BoolP("invert-sign", "i", false, "inverts the sign of the result.")
}
```

The `invert-sign` value is what will become the long flag. Users will be able to
include `--invert-sign` in their command call once we add this to our
subtractCmd. The `i` is the **short** flag. Users can do either the long flag,
or the short flag; they will mean the same thing.

The default value will be `false`, and the final string of text is just the
description of the flag's behavior.

Finally, this is a boolean flag, as denoted by the method call which is `BoolP`.
There are several other types of flags, such as StringP, IntP, etc. The `P` in
`BoolP` denotes that you also want to include a short flag, which is nice and
convenient for users. If you prefer not to include a short flag, just use
`Bool`, or whatever type of flag you want.

If you build and run the project now, you see that the `subtract` subcommand has
our new flag and its description.

```
$ go build . && ./math subtract -h
subtract integers

Usage:
  math subtract [flags]

Flags:
  -h, --help          help for subtract
  -i, --invert-sign   inverts the sign of the result.
```

That said, enabling the flag doesn't change anything, so we need to update our
`subtractCommandRunE` to use this value. Doing that is simple enough. Right
before we print things to the user, lets run our `subtract` function and then
invert the sign if the user requested it. The value of the flag is stored in our
`cmd` parameter.

```go
func subtractCommandRunE(cmd *cobra.Command, args []string) error {
    //   .. everything up here is unchanged ...

	result := subtract(values...)
	invert, _ := cmd.Flags().GetBool("invert-sign") // get the flag value!
	if invert {
		result = -result
	}

	fmt.Fprintln(cmd.OutOrStdout(), result)

	return nil
}
```