package main

import (
	"fmt"
)

func doSomething(x int) int {
	// does something
	return x * 5
}

// This task demonstrate the use of channels in a goroutine
func main() {

	ch := make(chan int)

	go func() {
		// use channel to get the returned value from the running fund in the go routine
		ch <- doSomething(7)

	}()

	fmt.Println(<-ch)

}
