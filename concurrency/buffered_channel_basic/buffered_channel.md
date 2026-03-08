# Buffered channel

Buffered channels block when no more capacity i.e. length == capacity. Changing range to > 10 will lead to a deadlock as no goroutines are alive reading from the channel
```shell
~/github_tangibleideas/hello-go/concurrency/buffered_channel_basic main*
❯ go run buffered_channel.go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        /Users/amit/github_tangibleideas/hello-go/concurrency/buffered_channel_basic/buffered_channel.go:10 +0x48
exit status 2
```

range can read all values from a channel, but requires a closed channel (eventually)

If there are no goroutines alive and writing to the channel and a goroutine still reading from an open channel go detects the deadlock and panics .

Commenting the close(ch) statement leads to 

```shell
❯ go run buffered_channel.go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        /Users/amit/github_tangibleideas/hello-go/concurrency/buffered_channel_basic/buffered_channel.go:18 +0x160
exit status 2
```