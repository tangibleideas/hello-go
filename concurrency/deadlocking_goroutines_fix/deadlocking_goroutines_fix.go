package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		toMain := 1
		ch1 <- toMain
		fromMain := <-ch2
		fmt.Println("goroutine:", "toMain:", toMain, "fromMain:", fromMain)
	}()

	// Read from ch1 first ( because func above accesses ch1 first)
	fromFunc := <-ch1

	toFunc := 2
	ch2 <- toFunc

	fmt.Println("goroutine:", "toFunc:", toFunc, "fromFunc:", fromFunc)
	// Allow func to recieve value from channel before main terminates the program
	time.Sleep(1 * time.Second)
}
