package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sum := 0
	for i := 1; i <= 5; i++ {
		random := rand.Intn(50-10) + 10
		fmt.Println(random)
		sum = sum + random
	}
	println("average is ", sum/5)
}
