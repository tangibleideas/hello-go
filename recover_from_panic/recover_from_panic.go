package main

import "fmt"

func main() {

	// uncomment myfunc() call
	// myfunc()
	handlePanic()
	fmt.Println("Main Done")

}

func myfunc() {
	ch := make(chan int, 1)
	fmt.Println("Created channel of length =", len(ch))
	close(ch)
	fmt.Println("Closed channel")

	// close a closed channel to cause panic
	close(ch)

}

func handlePanic() {
	defer func() {
		fmt.Println("Running deferred function")
		if r := recover(); r != nil {
			fmt.Println("Recoverd from panic", r)
		}
	}()

	fmt.Println("Running risky stuff")
	myfunc()
	fmt.Println("After myfunc")
}
