package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, _ := reader.ReadString('\n') // option 1
	fmt.Println(inputScore)
}
