# Deadlocking GoRoutines Fix

When two functions access multiple channels, they must be accessed in same order otherwise they create a deadlock

```shell
~/github_tangibleideas/hello-go main*
❯ go run concurrency/deadlocking_goroutines_fix/deadlocking_goroutines_fix.go
goroutine: toFunc: 2 fromFunc: 1
goroutine: toMain: 1 fromMain: 2
```
