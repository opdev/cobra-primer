# Implementing Logic

So now that we have a `sum` command, we can implement some logic. 

If you take a look at the `sumCmd`, you'll notice that it has a key called `Run`
with a scaffolded anonymous function that runs `fmt.Println("sum called")`.
That's what we saw in the last section when we ran `math sum`.

```go
var sumCmd = &cobra.Command{
	Use:   "sum",
    ...
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sum called")
	},
}
```

These `Run` functions are where you will implement the business logic for your
subcommand. The `Run` documentation is below:

```go
    // Run: Typically the actual work function. Most commands will only implement this.
	Run func(cmd *Command, args []string)
```

Effectively, the `Run` key in a `cobra.Command` struct just needs to be a
function that has this exact signature. By convention, I prefer to have these
functions defined (as opposed to being anonymous functions), so you might see me
do something like this:

```go
var sumCmd = &cobra.Command{
	Use:   "sum",
    ...
	Run: sumCommandRun,
}

func sumCommandRun(cmd *cobra.Command, args []string) {
		fmt.Println("sum called")
}
```

Subjectively, this makes the cobra command a little bit easier to read.

Let's replace this placeholder `println` with some logic. I'll write a sum
function that look something like this (there may already be sum functions, but
these are simple enough to write):

```go
func sum(values ...int) int {
	x := 0
	for _, v := range values {
		x += v
	}

	return x
}
```

You should be able to call this function with an arbitrary number of integer
values and get the sum of those values. Check it out on the [Go
Playground](https://go.dev/play/)

Our core logic is written, so now we just need to wire up the arguments that the
user provided to this function.
