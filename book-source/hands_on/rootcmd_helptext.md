# Adding Relevant Help Text

Our root command is currently returning scaffolded help text. That's not super
helpful - let's update that to something like this.

```go
var rootCmd = &cobra.Command{
	Use:   "math",
	Short: "Execute fun math functions",
}
```

I've removed the `Long` key from the `rootCmd`. For your actual projects,
provide something useful there, such as links to documentation, a longer
explanation of your goals with this tool, some examples - whatever you see fit.

Save, build and run the project, and you should see our new short text is
printed whenever I run it with no arguments:

```shell
$ go build -o math . && ./math
Execute fun math functions
```

Our help output is pretty bare right now, but it'll improve as we add
subcommands.