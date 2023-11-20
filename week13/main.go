package main

import "fmt"

func main() {
	var a []string
	fmt.Printf("%#v\n", a)
	fmt.Println(a, len(a), cap(a))
}
