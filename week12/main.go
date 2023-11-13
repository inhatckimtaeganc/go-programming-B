package main

import "fmt"

func main() {
	a := []string{"a", "b", "c", "d"}
	as := a[:2]
	as[1] = "Z"
	fmt.Println(a, as)

	b := [4]int{4, 3, 2, 1}
	bs := b[1:3]
	fmt.Println(b, bs)
}
