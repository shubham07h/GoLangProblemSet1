package main

import "fmt"

func main() {
	var a string
	fmt.Println("enter the value of first name")
	fmt.Scanln(&a)
	var b string
	fmt.Println("enter the value of last name")
	fmt.Scanln(&b)
	fmt.Println(a + " " + b)

}
