# The Cobra CLI

Using cobra in your project is as simple as importing `github.com/spf13/cobra`
where you need it, and then building out your command structure. If you're just
getting started using cobra, however, it may be beneficial to instead use
`cobra-cli`, which is a scaffolding tool for cobra applications.

We'll use that here, but just know that it's not a requirement. You can just as
easily start your CLI project by importing cobra and laying out your project as
you want. 

Follow the instructions [here](https://github.com/spf13/cobra#usage) for
installing the Cobra command line tool. The instructions indicate that you'll
install the "latest" version of `cobra-cli`, which at the time of this writing
appears to be v1.3.0. The commands you see here may differ slightly if you have
a different version, but that's okay.

Once you've installed `cobra-cli`, make sure it's in your `$PATH` by calling 

```
cobra-cli --help
```

If the you get a `command not found` error (or similar, then that implies that
your `$GOBIN` is not in your `$PATH`). Resolve that first before moving on. 

---

This book will build out a `math` binary that has several mathematical
operations encompassed as subcommands. Like I mentioned, it's a bit silly and
contrived, but we don't really want to implement a ton of logic here - we just
want to see what we can do with cobra.

The completed source code for the command built here will be available at
[https://github.com/opdev/cobra-primer/math](https://github.com/opdev/cobra-primer/math).