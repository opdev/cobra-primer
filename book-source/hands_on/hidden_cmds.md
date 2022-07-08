# Adding hidden commands

It's possible to add "hidden" commands, which are commands that do not show up
in help output but can be called. I don't have a great use case for it, but
either way, it's just a matter of making adding `cobra.Command.Hidden` and
setting it to true. I've scaffolded a subcommand `supersecretmath` that I've
marked hidden.

```
$ cobra-cli add supersecretmath
supersecretmath created at /Users/me/.go/src/github.com/opdev/cobra-primer/math
```

```go
// source: cmd/supersecretmath.go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// supersecretmathCmd represents the supersecretmath command
var supersecretmathCmd = &cobra.Command{
	Use:    "supersecretmath",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is where we do super secret math!")
	},
}

func init() {
	rootCmd.AddCommand(supersecretmathCmd)
}
```

We can't see this command in the help output!

```
$ go build . && ./math --help
Execute fun math functions

Usage:
  math [command]

Available Commands:
  arithmetic      basic arithmetic functions
  completion      Generate the autocompletion script for the specified shell
  help            Help about any command

Flags:
  -h, --help     help for math
  -t, --toggle   Help message for toggle

Use "math [command] --help" for more information about a command.
```

But we can certainly call it without a problem:

```
$ go build . && ./math supersecretmath
This is where we do super secret math!
```