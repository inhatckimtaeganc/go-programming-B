package main

import "fmt"

func main() {
	var s []int
	s = make([]int, 5)

	for _, value := range s {
		fmt.Println(value)
	}
}
