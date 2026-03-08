package main

import "fmt"

func main() {
	ch := make(chan int, 10)

	// Push values into the channel
	for i := range 11 {
		ch <- i
	}

	// Must be closed or the range can not detect that the channel is closed and blocks go knows that there is no goroutine alive and it's a deadlock (error in readme)
	close(ch)

	// Read all values from the channel
	for v := range ch {
		fmt.Println("ch : capacity =", cap(ch), "length =", len(ch))
		fmt.Println("read value=", v)
	}

}
