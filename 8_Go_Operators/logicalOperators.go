package main

import "fmt"

func main() {
	x := 7
	y := 10

	// &&  (Logical AND)
	// Returns true ONLY if **both** conditions are true
	condAnd := x > 5 && y > 5 // true && true => true
	fmt.Println("x > 5 && y > 5:", condAnd)

	condAnd2 := x > 5 && y < 5 // true && false => false
	fmt.Println("x > 5 && y < 5:", condAnd2)

	// ||  (Logical OR)
	// Returns true if **at least one** condition is true
	condOr := x > 5 || y < 5 // true || false => true
	fmt.Println("x > 5 || y < 5:", condOr)

	condOr2 := x < 5 || y < 5 // false || false => false
	fmt.Println("x < 5 || y < 5:", condOr2)

	// !  (Logical NOT)
	// Reverses the result — true becomes false, false becomes true
	condNot := !(x > 5 && y > 5) // !(true && true) => !true => false
	fmt.Println("!(x > 5 && y > 5):", condNot)

	condNot2 := !(x < 5 || y < 5) // !(false || false) => !false => true
	fmt.Println("!(x < 5 || y < 5):", condNot2)
}
