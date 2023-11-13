package main

import "fmt"

func main() {
	// var s []int
	// s = make([]int, 5)

	// s := make([]int, 5) // 단축연산자 사용, 변수 선언과 동시에 메모리 할당, 원소들은 제로 값으로 초기화

	s := []int{0, 0, 0, 0, 0} // 단축연산자 사용, 변수 선언과 동시에 메모리 할당 그리고 원소 초기화

	s[4] = 100
	s[0] = 7
	s[1] = 91

	for _, value := range s {
		fmt.Println(value)
	}
}
