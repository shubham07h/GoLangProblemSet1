package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string = "123"

	num, err := strconv.ParseInt(a, 6, 12)
	if err != nil {

	}
	fmt.Println(num)
	println(num * 10)

}
