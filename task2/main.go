package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go printNumbers(i, &wg)
	}
	letters := "abcdefghijklmnopqrstuvwxyz"
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go printLetters(string(letters[i]), &wg)
	}

	wg.Wait()
}

func printNumbers(number int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(number)
}

func printLetters(letter string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(letter)
}
