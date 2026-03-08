package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		toMain := 1
		ch1 <- toMain
		fromMain := <-ch2
		fmt.Println("goroutine:", "toMain:", toMain, "fromMain:", fromMain)
	}()

	toFunc := 2
	ch2 <- toFunc
	fromFunc := <-ch1

	fmt.Println("goroutine:", "toFunc:", toFunc, "fromFunc:", fromFunc)
}
