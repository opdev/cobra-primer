# Enforcing Expectations on Arguments

What happens if we don't provide any arguments to our `math sum` subcommand?

```shell
# ./math sum
0
```

I supposed that's technically correct! We didn't provide arguments, so the sum
of 0 is... well 0. But instead, let's make sure the user provides values, or
otherwise print the help output. One way to do that is using the
`cobra.Command.Args` key. This key specifies that its type is `PositionalArgs`.
If we look through cobra's documentation, we see that type defined as a function
with this signature:

```
type PositionalArgs func(cmd *Command, args []string) error
```

Check it out in the documentation
[here](https://pkg.go.dev/github.com/spf13/cobra#PositionalArgs).

That means we can write a function that fits that signature, and pass it to
`sumCmd`'s Args key. Since we've got `args` as a parameter here as well, this is
actually pretty easy to write. But Cobra actually makes some pre-defined
functions available to us, and one of those does exactly what we want:

```go
// MinimumNArgs returns an error if there is not at least N args. 
func MinimumNArgs(n int) PositionalArgs`
```
[Doc](https://pkg.go.dev/github.com/spf13/cobra#MinimumNArgs)

We can use that to enforce the expectation that we have at least 1 argument by
passing this function to the `Args` key in our `sumCmd`:

```go
// sumCmd represents the sum command
var sumCmd = &cobra.Command{
	Use:   "sum",
	Args:  cobra.MinimumNArgs(1),   // Add it here
	Short: "A brief description of your command",
    // ...
}
```

Now the command properly indicates that we didn't provide enough arguments when
called:

```shell
# go build -o math . && ./math sum 
Error: requires at least 1 arg(s), only received 0
Usage:
  math sum [flags]

Flags:
  -h, --help   help for sum
```

There are several other PositionalArg (or `Args` compatible) functions in the
cobra library, but since we know the function signature, we can build our own.

In addition the cobra library also lets us enforce multiple requirements on our
arguments as well by using the `cobra.MatchAll` function
[doc](https://pkg.go.dev/github.com/spf13/cobra#MatchAll).