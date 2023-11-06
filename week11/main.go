package main

import "fmt"

func main() {
	// var primes [3]int
	// primes[0] = 2
	// primes[1] = 3
	// primes[2] = 5
	// fmt.Println(primes, primes[1])

	// var primes [3]int = [3]int{2, 3, 5} // 배열 리터럴로 primes 배열 초기화
	// fmt.Println(primes, primes[1])

	primes := [3]int{2, 3, 5} // 배열 리터럴로 primes 배열 초기화
	fmt.Println(primes, primes[1])

	test := [5]bool{true, true, true} // 초기화 하지 않은 배열 원소의 제로 값은 해당 원소의 타입의 제로값으로 결정된다
	fmt.Println(test[3])

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(primes[i])
	// }

	//fmt.Println(primes[3]) // 컴파일 에러. invalid argument: index 3 out of bounds [0:3]

	// i := 0
	////for i < 4 { // 실행시간 에러. panic: runtime error: index out of range [3] with length 3
	// for i < len(primes) {
	// 	fmt.Println(primes[i])
	// 	i++
	// }

	// for idx, prime := range primes {
	// 	fmt.Println(idx, prime)
	// }

	//for prime := range primes {  // 소수 값을 배열에서 꺼내려고 하지만 인덱스가 출력
	//for idx, prime := range primes { // 선언하고 변수를 사용 안함. 컴파일 에러
	for _, prime := range primes { // 인덱스 사용 안함
		fmt.Println(prime)
	}
}
