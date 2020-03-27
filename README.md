# Remove-k8s-ns-finalizers

*Note: Due to not having a Windows/MacOS system, this tool has been untested on those platforms!*

*Note 2: Only use this tool if you know what you're doing! Removing finalizers can have severe consequences!*

Sometimes a Kubernetes namespaces is 'stuck' in the Terminating phase. This is in most cases, because there is some finalizer left that isn't handled.

Via some manual API calls, this can be fixed, see for instance:
- https://stackoverflow.com/questions/52369247/namespace-stuck-as-terminating-how-do-i-remove-it
- https://medium.com/@clouddev.guru/how-to-fix-kubernetes-namespace-deleting-stuck-in-terminating-state-5ed75792647e

This simple tool automates this step with a single command.


## Usage

Make sure you're logged into the right cluster and switched to the right k8s context.

After that, just execute this command with the namespace you want to finalize/terminate as its argument:

```
$ ./remove-k8s-ns-finalizers your-namespace
```


## How to build

Make sure you have a recent version of Go (tested with 1.14), after that just call make:

```
$ make
```
