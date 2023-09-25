package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	a := 7
	fmt.Println(a, reflect.TypeOf(a))

	var b float32 = 8.34
	fmt.Println(a * int(b))
	fmt.Println(float32(a) > b)

	var c string
	var d int
	var e bool
	var f float64
	fmt.Println(c, d, e, f)

	fmt.Println('Z', '2', '3', '4')
	fmt.Println(math.Floor(2.17), math.Ceil(2.17), math.Sqrt(16))
	fmt.Println(strings.Title("hello world"))
}
