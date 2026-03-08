# Recover from panic

Wrap the operation that may panic into a function and defer call to recover().
recover() stops panic and returns panic message 

Uncomment call to myfunc() in main, it'll cause
```shell
~/github_tangibleideas/hello-go/recover_from_panic main*
❯ go run recover_from_panic.go
Created channel of length = 0
Closed channel
panic: close of closed channel

goroutine 1 [running]:
main.myfunc()
        /Users/amit/github_tangibleideas/hello-go/recover_from_panic/recover_from_panic.go:21 +0xc0
main.main()
        /Users/amit/github_tangibleideas/hello-go/recover_from_panic/recover_from_panic.go:8 +0x1c
exit status 2
```

If wrapper func used 

```shell
~/github_tangibleideas/hello-go/recover_from_panic main*
❯ go run recover_from_panic.go
Running risky stuff
Created channel of length = 0
Closed channel
Running deferred function
Recoverd from panic close of closed channel
Main Done
```