package main

import (
	"fmt"
)

func main() {
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	countries := [...]string{"Kenya", "Libya", "Ethiopia"}
	fmt.Print(cars, countries)
}
