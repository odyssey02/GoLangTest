package main

import (
	"fmt"
	"time"
)

func main() {
	//var now time.Time = time.Now()
	var now time.Time
	now = time.Now()
	var year = now.Year()
	month := now.Month()
	fmt.Println(year, month, now.Day(), now.Hour(), now.Minute(), now.Second())
}
