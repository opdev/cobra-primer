# Wiring up Arguments

At this point, we have our business logic (the `sum(...)` function) written and
working. We just need to replace the placeholder code in our `sumCommandRun` so
that it uses the `sum` function. We need to get the user's arguments passed over
to our `sum` function so that everything works.

If you look at the function signature required for the `cobra.Command.Run`
struct key,  and our `sumCommandRun` function, you'll see that `args` is a
parameter we can use, and it contains the arguments passed in by the user,
without flags or the subcommand structure.

So for example, if the user ran `math sum 1 2 3`, then we'd expect `args` to be
`[]string{"1", "2", "3"}`.

With that in mind, the problem we have with `args` is that it's a `[]string`.
Let's convert that over to a `[]int` which is what our `sum` function uses. The
logic to do this isn't important, but helps us complete this example.

```go
func sumCommandRun(cmd *cobra.Command, args []string) {
	// convert args which is []string to []int
	values := make([]int, len(args))
	for i, v := range args {
		vAsInt, _ := strconv.Atoi(v)
		values[i] = vAsInt
	}

	fmt.Println(sum(values...))
}
```

Then build and run the `sum` subcommand:

```shell
# go build -o math . && ./math sum 2 3 4
9
```

Everything works!

So the `args` parameter that we have to work with here in our `Run` function
contains all of the arguments that are passed in by the user to this subcommand.
Feel free to `fmt.Println([]args)`, and then run `math sum` with random values
to see what gets printed.