package main

import "fmt"

func main() {
	// var s []int
	// s = make([]int, 5)

	// s := make([]int, 5) //단축연산자 사용 초기화

	s := []int{0, 0, 0, 0, 0} //단축연산자, 슬라이스리터럴 사용 초기화

	s[4] = 100
	s[0] = 7
	s[1] = 91

	for _, value := range s {
		fmt.Println(value)
	}
}
