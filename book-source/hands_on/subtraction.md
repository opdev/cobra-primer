# Adding another subcommand

Let's break out `cobra-cli` and add a `subtract` subcommand! Run this from the
base of your repository.

```
$ cobra-cli add subtract
subtract created at /Users/me/.go/src/github.com/opdev/cobra-primer/math
```

Now we've got a `cmd/subtract.go` file:

```
.
├── LICENSE
├── cmd
│   ├── root.go
│   ├── subtract.go
│   └── sum.go
├── go.mod
├── go.sum
├── main.go
└── math
```

Modify the `Long` and `Short` descriptions as you see fit. As a practice, try the following actions:


- Make the command accept only 2 positional arguments, e.g. `math subtract 2 3`
- Swap out the `Run` function with a standalone `RunE` function.
- Convert the `args` values to integers, returning errors if encountered (copy
  this from the `sum` command, or better yet, make it its own function and reuse
  it here).

When done, the `math subtract` function should work like this:

```
$ ./math subtract 2 3
-1

$ ./math subtract 2 3 4
Error: accepts 2 arg(s), received 3
Usage:
  math subtract [flags]

Flags:
  -h, --help   help for subtract
```
