package main

import (
	"fmt"
	"sync"
)

func printMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done() //mark gproutine as complete
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2) //We will launch 2 goroutines

	go printMessage("Hello", &wg)
	go printMessage("World", &wg)

	wg.Wait() //Wait for all goroutines to complete
	fmt.Println("All goroutines complete")
}
