package main

import "fmt"

func main() {
	a := make([]string, 4, 5) // type, length, capacity
	a[0] = "a"
	a[2] = "c"
	a[3] = "d"
	as := a[0:2]
	as[1] = "Z"
	c := append(a, "y")

	fmt.Println(a, len(a), cap(a))
	fmt.Println(c, len(c), cap(c))
}
