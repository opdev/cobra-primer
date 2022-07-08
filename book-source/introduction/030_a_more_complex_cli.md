# A more complex interface


Let's take a look at a typical `kubectl` invocation:

```shell
kubectl apply -f mypod.yaml
```

This commandline interface looks a little different in that now we've got this
structure here:

```txt
COMMAND SUBCOMMAND  SHORTFLAG   ARGUMENT
kubectl apply       -f          mypod.yaml
```
So here we see this idea of a **subcommand**.

The `kubectl` command itself takes on a more organizational role, in that it
almost _defines_ the context for a subcommand ("We're working with a kubernetes
cluster").

Running `kubectl` by itself generally does not take any action, other
than displaying help text. But while it does not take any action, it may itself
have some of its own flags that apply to all subcommands (see `kubectl
options`).

The actual action we're taking here is described by the subcommand: `apply`
("we're going to take some data and apply it to the cluster").

And if you're familiar with `kubectl`, you'll know that it gives us many more
commands to interact with our cluster (such as `get`, `describe`, `create`,
etc). And each subcommand may have its own flags, on top of those already built
into the top-level `kubectl command`. 

Building this kind of commandline interface with just the **flag** package can
be complex. The **flag** package helps us extract short and long flags out of
the arguments passed into our binary by the user, but we need to build in these
relationships such that each subcommand may have its own flags, and the base
command itself has its own flags. Before you know it, you've written more logic
to handle the relationships between your commands than your actual business
logic.

This is the problem the [Cobra](cobra.dev) library solves for us.
