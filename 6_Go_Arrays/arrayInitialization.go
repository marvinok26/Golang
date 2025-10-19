package main

import (
	"fmt"
)

// If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.
// The default value for int is 0, and the default value for string is "".
func main() {
	arr1 := [5]int{}              //not initialized
	arr2 := [5]int{1, 2}          //partially initialized
	arr3 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}
