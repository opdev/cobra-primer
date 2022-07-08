# Returning an error

In the last section, we wired up our user's arguments and passed it to our `sum`
function. What happens if the user passes in a string value? What about a
decimal?

```shell
# ./math sum 1 2 3 foo
6

# ./math sum 1 2 3 2.1
6
```

Our sum function isn't even batting an eye! It's just completely ignoring the
string `foo`, and the decimal `2.1` value.  That's because we ignored our error
when we converted from `[]string` to `[]int`. Take a look at the line with block
`[1]`.:

```go
// ...
    for i, v := range args {
            vAsInt, _ := strconv.Atoi(v)       // [1]
            values[i] = vAsInt
    }
// ...
```

Here we use the `strconv` librarie's `Atoi` function to convert the string value
to an integer, and then we disregard the second return value which is an error.
Ideally, we want to return that error, but we have a bit of a problem. Our
`sumCommandRun` function doesn't return an error, but that's easy enough to fix:

```go
func sumCommandRun(cmd *cobra.Command, args []string) error { // return an error
	// convert args which is []string to []int
	values := make([]int, len(args))
	for i, v := range args {
		vAsInt, err := strconv.Atoi(v)
        if err != nil {     // new code!
            return fmt.Errorf("you provide a value that was not an integer: %s", v)
        }
		values[i] = vAsInt
	}

	fmt.Println(sum(values...))

    return nil // all went well, return no error
}
```

As soon as we reconfigure our function, our `sumCmd` should show an error that
reads:

```text
cannot use sumCommandRun
    (value of type func(cmd *cobra.Command, args []string) error)
as
    func(cmd *cobra.Command, args []string) value in struct literal
```

This is because the `cobra.Command.Run` key enforces a specific function
signature that matched what we were using previously. If we want to return an
error in our function, we can do so by instead assigning our function to `RunE`.
Its function signature is identical, but it returns an error.

```Go
//...
var sumCmd = &cobra.Command{
	Use:   "sum",
    //...

    //Run: sumCommandRun        // this is what we used before
	RunE: sumCommandRunE,        // and replaced it with this.
}
/...
```

By convention, I've also renamed our `sumCommandRun` to `sumCommandRunE` to make
it match the command struct key to which it applies.

The project should be happy, and you should see an error returned when the user
provides non-integer values.

```
$ go build -o math . && ./math sum 1 2 3 foo
Error: you provide a value that was not an integer: foo
Usage:
  math sum [flags]

Flags:
  -h, --help   help for sum
```

And since we've returned an error, the help output is provided to the user. 