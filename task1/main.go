package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		go printNumbers(i)
	}
	letters := "abcdefghijklmnopqrstuvwxyz"
	for i := 1; i < 10; i++ {
		go printLetters(string(letters[i]))
	}

	time.Sleep(2 * time.Second)
}

func printNumbers(number int) {
	fmt.Println(number)
}

func printLetters(letter string) {
	fmt.Println(letter)
}
