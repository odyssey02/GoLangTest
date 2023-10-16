package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("hello")
}

func main() {
	var a int = 10
	n := 3.0
	p := 282
	g := 1
	fmt.Println("hi", float64(p+g)/3.0, n, a)
	sayHello()
}
