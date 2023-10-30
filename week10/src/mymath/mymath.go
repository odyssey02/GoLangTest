package mymath

//func myPower(base int, exponent int) int {
func MyPower(base int, exponent int) int { //외부에 함수를 공개하려면 이름 첫글자가 대문자여야함
	result := 1
	for i := 1; i <= exponent; i++ {
		result *= base
	}
	return result
}

func MyAbs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
