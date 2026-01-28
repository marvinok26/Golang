package main

import "fmt"

// basic example of "send something to a goroutine"
func main() {
	ch := make(chan string)

	go func() {
		ch <- "Goroutine says Hello"
	}()

	msg := <-ch
	fmt.Println(msg)
}
