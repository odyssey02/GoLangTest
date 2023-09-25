package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guess number game!")
	answer := rand.Intn(100) + 1 //1~100
	fmt.Println(answer)

	for i := 0; i < 10; i++ {
		fmt.Print("Input number : ")
		reader := bufio.NewReader(os.Stdin)
		inputNumberString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputGuessString := strings.TrimSpace(inputNumberString)
		inputGuess, err := strconv.Atoi(inputGuessString)
		if err != nil {
			log.Fatal(err)
		}
		if inputGuess < answer {
			fmt.Println("Your guess number is lower than answer!") //answer is higher
		} else if inputGuess > answer {
			fmt.Println("Your guess number is higher than answer!") //answer is lower
		} else {
			fmt.Println("Your guess number is correct!")
			break
		}
	}
}
