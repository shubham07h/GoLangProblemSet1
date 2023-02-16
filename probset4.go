package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("enter the place")
	var place string
	fmt.Scanln(&place)
	var capital string = strings.ToUpper(place)
	fmt.Println("value in uppercase is", capital)
	var small string = strings.ToLower(capital)
	fmt.Println("value in lowercase is", small)

}
