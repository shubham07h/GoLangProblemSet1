package main

import (
	"fmt"
)

func main() {
	var diameter float32
	fmt.Println("Enter value of diameter")
	fmt.Scanln(&diameter)
	var radius float32 = diameter / 2
	var perimeter float32 = 2 * 3.14 * radius
	var area float32 = 3.14 * radius * radius
	fmt.Println("value of perimeter", perimeter)
	fmt.Println("value of area", area)

}
