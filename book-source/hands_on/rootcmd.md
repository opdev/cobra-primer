# The Root Command

Inside the `cmd` directory ("package") that `cobra-cli` scaffolded for you
should be a file called `root.go`. Let's look at some key parts of this file
(I've truncated a few sections with ellipsis, `...`, so that it appears cleaner in this book.)

```go
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "math",
	Short: "A brief description of your application",
	Long: `...`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}
```

Here we define the variable `rootCmd` at the global scope to be a
`*cobra.Command`. the `cobra.Command` struct is the building block of
cobra-based applications, and is how we define almost everything we need
(metadata, logic, etc.) to execute our business logic. Every command and
subcommand will be defined as one of these `cobra.Command`s. We'll use these to
build out our tree of subcommands.

As we saw, our `main()` function calls the `Execute()` function in this package,
which itself just wraps our `rootCmd`. 

```go
// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
```

As the scaffolded comment suggests, this is called by `main()`, and only ever
needs to happen here.

---

As mentioned in the `kubectl` context, the root command of a non-trivial CLI
application is typically organizational, and typically doesn't have any logic
associated with it. That way, when a user runs our command with no flags,
subcommands, args, or otherwise, they get the help output in return to help
guide them.
