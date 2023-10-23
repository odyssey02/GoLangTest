package main

import "fmt"

func double(n *int) {
	*n *= 2
}
func main() {
	var a int = 5
	double(&a)
	fmt.Printf("%d\n", a)
}
