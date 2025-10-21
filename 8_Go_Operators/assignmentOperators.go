package main

import "fmt"

func main() {
	// ASSIGNMENT OPERATORS
	x := 5
	fmt.Println("Start: x = 5")

	x = 5 // simple assignment
	fmt.Println("x = 5     →", x)

	x += 3 // x = x + 3
	fmt.Println("x += 3    →", x)

	x -= 3 // x = x - 3
	fmt.Println("x -= 3    →", x)

	x *= 3 // x = x * 3
	fmt.Println("x *= 3    →", x)

	x /= 3 // x = x / 3
	fmt.Println("x /= 3    →", x)

	x %= 3 // x = x % 3
	fmt.Println("x %= 3    →", x)

	// BITWISE ASSIGNMENT OPERATORS
	fmt.Println("\nNow the bitwise assignment operators with x = 5 and y = 3")

	// Reset x
	x = 5
	fmt.Println("\nBinary: 5 = 0101, 3 = 0011")

	// &= AND → keep only bits that are 1 in BOTH numbers
	x &= 3
	fmt.Println("x &= 3    →", x, " (AND)")

	// Reset x
	x = 5

	// |= OR → keep any bit that is 1 in EITHER number
	x |= 3
	fmt.Println("x |= 3    →", x, " (OR)")

	// Reset x
	x = 5

	// ^= XOR → keep bits ONLY where the bits are different
	x ^= 3
	fmt.Println("x ^= 3    →", x, " (XOR)")

	// Reset x
	x = 5

	// >>= RIGHT SHIFT → move bits to the right (divide by 2)
	x >>= 1
	fmt.Println("x >>= 1   →", x, " (Right Shift)")

	// Reset x
	x = 5

	// <<= LEFT SHIFT → move bits to the left (multiply by 2)
	x <<= 1
	fmt.Println("x <<= 1   →", x, " (Left Shift)")
}
