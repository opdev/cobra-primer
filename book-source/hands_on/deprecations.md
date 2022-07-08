# Marking commands deprecated

Say you need to deprecate the `sum` command in favor of an `addition` command.
In this case, you simply add a `cobra.Command.Deprecated` key with a string
indicating the message you want printed to the user.

```go
// source: cmd/sum.go
var sumCmd = &cobra.Command{
    // ...nothing else changed...
	Deprecated: `This command will be replaced by the "addition" command in the next release`,
    // ...nothing else changed...

	RunE: sumCommandRunE,
}
```

And the message is passed to the user when this command is called.

```
$ go build . && ./math arithmetic sum --help
Command "sum" is deprecated, This command will be replaced by the "addition" command in the next release
Given an arbitrary number of integer arguments,

this will return the sum of all values.

Usage:
  math arithmetic sum [flags]

Aliases:
  sum, total

Flags:
  -h, --help   help for sum

13:20:44 ~/.go/src/github.com/opdev/cobra-primer/math
```