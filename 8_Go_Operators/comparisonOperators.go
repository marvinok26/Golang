package main

import "fmt"

func main() {
	x := 10
	y := 20

	// ==  Equal to
	// Checks whether x is equal to y
	isEqual := x == y
	fmt.Println("x == y:", isEqual) // false

	// !=  Not equal
	// Checks whether x is NOT equal to y
	isNotEqual := x != y
	fmt.Println("x != y:", isNotEqual) // true

	// >  Greater than
	// Checks if x is strictly greater than y
	isGreater := x > y
	fmt.Println("x > y:", isGreater) // false

	// <  Less than
	// Checks if x is strictly less than y
	isLess := x < y
	fmt.Println("x < y:", isLess) // true

	// >=  Greater than or equal to
	// Checks if x is greater than OR equal to y
	isGreaterOrEqual := x >= y
	fmt.Println("x >= y:", isGreaterOrEqual) // false

	// <=  Less than or equal to
	// Checks if x is less than OR equal to y
	isLessOrEqual := x <= y
	fmt.Println("x <= y:", isLessOrEqual) // true
}
