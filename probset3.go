package main

import "fmt"

func main() {
	fmt.Println("enter value in celcius")
	var celcius float32
	fmt.Scanln(&celcius)
	var fahrenhit float32 = (celcius*9)/5 + 32
	fmt.Println("value of fahrenhit is ", fahrenhit)

}
