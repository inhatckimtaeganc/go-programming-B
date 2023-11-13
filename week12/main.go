package main

import "fmt"

func main() {
	s := []int{0, 0, 0, 0, 0} // 단축 연산자 및 슬라이스 리터럴 사용, 변수 선언과 동시에 메모리 할당 그리고 원소 초기화

	s[4] = 100
	s[0] = 7
	s[1] = 91

	for _, value := range s {
		fmt.Println(value)
	}

	copyS := s[1:4]
	for _, value := range copyS {
		fmt.Println(value)
	}

	test := [3]string{"inha", "go", "student"}
	testS := test[:2] // testS := test[0:2]
	fmt.Println(test, len(test))
	fmt.Println(testS, len(testS))
}
