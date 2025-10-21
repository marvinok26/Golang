package main

import "fmt"

func main() {
	x := 6 // binary: 110
	y := 3 // binary: 011

	// Bitwise AND (&)
	// 110 (6)
	// 011 (3)
	// ---- AND
	// 010 = 2
	fmt.Println("x & y =", x&y)

	// Bitwise OR (|)
	// 110 (6)
	// 011 (3)
	// ---- OR
	// 111 = 7
	fmt.Println("x | y =", x|y)

	// Bitwise XOR (^)
	// 110 (6)
	// 011 (3)
	// ---- XOR (1 only if bits are different)
	// 101 = 5
	fmt.Println("x ^ y =", x^y)

	// Left Shift <<
	// 6 is 110 in binary
	// 6 << 1 shifts left: 110 -> 1100 = 12
	// Equivalent to 6 * 2¹
	fmt.Println("x << 1 =", x<<1)

	// Right Shift >>
	// 6 is 110 in binary
	// 6 >> 1 shifts right: 110 -> 11 = 3
	// Equivalent to 6 / 2¹
	fmt.Println("x >> 1 =", x>>1)
}
