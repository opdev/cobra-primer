# Adding an alias

As your start to develop your tool, there may be cases where you want to `alias` a given subcommand. This is very easy to do, as there is an `Alias` key in the `cobra.Command` struct. Let's say that `total` is an alias of `sum`, such that a user can call `math total` and get the same logic.


```go
var sumCmd = &cobra.Command{
	Use:     "sum",
	Aliases: []string{"total"},         // total!
    // ... nothing else changed
}
```

Now run the command with `total` instead of `sum`, and see the same logic applied.

```shell
# go build -o math . && ./math total 1 2 3
6
```

This isn't a perfect user experience. You'll notice that `total` is not shown in the subcommand list:

```
./math -h
Execute fun math functions

Usage:
  math [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  sum         A brief description of your command

Flags:
  -h, --help     help for math
  -t, --toggle   Help message for toggle

Use "math [command] --help" for more information about a command.
```

It is shown in the help output for `sum`, however:

```
./math sum -h
Given an arbitrary number of integer arguments,

this will return the sum of all values.

Usage:
  math sum [flags]

Aliases:
  sum, total

Flags:
  -h, --help   help for sum
```

To that end, aliases are mostly helpful in cases where you have shorthand names
for common functions. An example might be `cp` for `copy`, `mv` for `move`, and
`rm` for `remove`.