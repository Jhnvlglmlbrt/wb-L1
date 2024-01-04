package main

import (
	"fmt"
	"sync"
)

func main() {
	inputCh := make(chan int)
	outputCh := make(chan int)

	var wg sync.WaitGroup
	numbers := []int{2, 4, 6, 8, 10}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(inputCh)

		for _, num := range numbers {
			inputCh <- num
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(outputCh)

		for num := range inputCh {
			result := num * 2
			outputCh <- result
		}
	}()

	for result := range outputCh {
		fmt.Println(result)
	}

	wg.Wait()
}
