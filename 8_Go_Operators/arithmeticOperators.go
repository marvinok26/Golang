package main

import "fmt"

func main() {
	// Declare two integers
	x := 10
	y := 3

	fmt.Println("x =", x, ", y =", y)

	// Basic arithmetic operations
	fmt.Println("Addition (x + y):", x+y)
	fmt.Println("Subtraction (x - y):", x-y)
	fmt.Println("Multiplication (x * y):", x*y)
	fmt.Println("Division (x / y):", x/y) // integer division
	fmt.Println("Modulus (x % y):", x%y)  // remainder

	// Increment and decrement
	x++ // increases x by 1
	fmt.Println("Incremented x (x++):", x)

	x-- // decreases x by 1
	fmt.Println("Decremented x (x--):", x)
}
