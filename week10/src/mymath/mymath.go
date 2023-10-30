package mymath

//func myPower(base int, exponent int) int {
func MyPower(base int, exponent int) int { // 외부에 함수를 공개하려면 반드시 함수 이름의 첫 글자를 대문자로 작성한다.
	result := 1
	for i := 1; i <= exponent; i++ {
		result = result * base
	}
	return result
}

func MyAbs(number int) int {
	if number < 0 { // 음수면
		return number * -1
	}
	return number // 양수면
}
