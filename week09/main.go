package main

import "fmt"

func main() {
	a := 10 // var a int = 10
	var pa *int
	pa = &a

	fmt.Println(a, *pa)
	fmt.Println(&a, pa)
	fmt.Printf("%T %T %T %T\n", a, *pa, &a, pa)
	fmt.Println(&pa)
}
