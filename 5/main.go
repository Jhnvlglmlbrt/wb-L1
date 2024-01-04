package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <numSeconds>")
		return
	}

	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("invalid command line argument")
		return
	}

	var wg sync.WaitGroup
	timer := time.After(time.Duration(N) * time.Second)

	mainChanel := make(chan int)

	// consumer
	wg.Add(1)
	go consumer(mainChanel, &wg)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-timer:
				fmt.Println("Timer finished.")
				close(mainChanel)
				return
			default:
				mainChanel <- rand.Intn(100)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	wg.Wait()
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Читаем из канала до закрытия
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("Channel closed. Exiting readFromChannel.")
			return
		}
		fmt.Printf("Received value: %d\n", value)
	}
}
