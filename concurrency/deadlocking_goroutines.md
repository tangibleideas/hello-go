# Deadlocking GoRoutines

When two functions access multiple channels, they must be accessed in same order otherwise they create a deadlock

```shell
❯ go run ./concurrency/deadlocking_goroutines.go 
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        /Users/amit/github_tangibleideas/hello-go/concurrency/deadlocking_goroutines.go:17 +0xa0

goroutine 35 [chan send]:
main.main.func1()
        /Users/amit/github_tangibleideas/hello-go/concurrency/deadlocking_goroutines.go:11 +0x34
created by main.main in goroutine 1
        /Users/amit/github_tangibleideas/hello-go/concurrency/deadlocking_goroutines.go:9 +0x8c
exit status 2
```
