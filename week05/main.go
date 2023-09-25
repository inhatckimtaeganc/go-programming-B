package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Guess number game!")
	answer := rand.Intn(100) + 1 // 1 ~ 100
	fmt.Println(answer)
}
