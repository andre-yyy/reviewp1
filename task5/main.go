package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	oddChannel := make(chan int)
	evenChannel := make(chan int)

	wg.Add(1)

	go sendNumber(oddChannel, evenChannel, &wg)

	go func() {
		for {
			select {
			case odd := <-oddChannel:
				fmt.Println("Received an odd number:", odd)
			case even := <-evenChannel:
				fmt.Println("Received an odd number:", even)
			}
		}
	}()

	wg.Wait()

}

func sendNumber(odd chan int, even chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(odd)
	defer close(even)
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			odd <- i
		} else {
			even <- i
		}
	}
}
