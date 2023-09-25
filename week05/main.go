package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScoreString, err := reader.ReadString('\n') //option 2
	if err != nil {
		log.Fatal(err)
	}
	inputScoreString = strings.TrimSpace(inputScoreString)      //remove space
	inputScore, err := strconv.ParseFloat(inputScoreString, 32) //string to 32bit float
	var grade string
	if inputScore >= 90 { //mismatched types string and untyped int
		grade = "A grade!"
	} else {
		grade = "under A grade..."
	}
	fmt.Println("Your got", grade)
}
