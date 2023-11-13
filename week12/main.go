package main

import "fmt"

func main() {
	s := []int{0, 0, 0, 0, 0} //단축연산자, 슬라이스리터럴 사용 초기화

	s[4] = 100
	s[0] = 7
	s[1] = 91

	for _, value := range s {
		fmt.Println(value)
	}

	copys := s[1:4]
	for _, value := range copys {
		fmt.Println(value)
	}

	test := [3]string{"wade", "finn", "jake"}
	testS := test[:2]
	testS2 := test[1:]

	// testS2[0] = "python"
	// test[1] = "python"
	testS[1] = "python"

	fmt.Println(test, len(test))
	fmt.Println(testS, len(testS))
	fmt.Println(testS2, len(testS2))

}
