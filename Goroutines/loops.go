package main

import "fmt"

func printNumbers(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println("Number", i)
	}
}

func checkNumber(n int) {
	if n%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func main() {
	printNumbers(7)
	checkNumber(4)
}
