package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	taskChannel := make(chan int, 5)

	wg.Add(2)

	go produce(taskChannel, &wg)
	go consume(taskChannel, &wg)

	wg.Wait()
}

func produce(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(c)

	for i := 1; i <= 10; i++ {
		c <- i
	}
}

func consume(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range c {
		fmt.Println(task)
	}
}
