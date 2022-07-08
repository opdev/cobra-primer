# The anatomy of a commandline interface

Let's take a look at a typical commandline interface. Say we invoke the `cat`
command:

```shell
cat -n myfile.txt
```

We could describe this as:
- `cat` is our **command**.
- `-n` is a **short flag**.
- `myfile.txt` is our **argument**.

Building something like this call is fairly simple in golang with just the
standard library. Building our `main.go` (or equivalent) file would produce a
binary on disk, which acts as our **command**.

We would just need to reach into the standard library for the
[flag](https://pkg.go.dev/flag) package, which gives us the ability to register
and parse **flags**.

Finally, we would just need to review the argument list to grab the final value of
`myfile.txt` that the user provided, and take action.

But what if we want something more complex?

_For reference, a naive implementation of `cat` with the `-n` flag using the
**flag** package is [here](./appendix/gocat.md)._