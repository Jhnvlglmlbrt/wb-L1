package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <numWorkers>")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid number of workers:", os.Args[1])
		return
	}

	mainChanel := make(chan int)
	var wg sync.WaitGroup

	// Горутина для ожидания сигналов и закрытия канала
	info := make(chan os.Signal, 1)
	signal.Notify(info, syscall.SIGINT, os.Interrupt)
	go func() {
		<-info
		close(mainChanel)
	}()

	// Горутины воркеров
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(mainChanel, &wg, i)
	}

	// Горутина для постоянной записи в канал
	go func() {
		defer wg.Done()
		for {
			select {
			case <-mainChanel:
				return // Выйти, если канал закрыт
			default:
				mainChanel <- rand.Intn(100)
				time.Sleep(2 * time.Second)
			}
		}
	}()

	wg.Wait()
	fmt.Println("All workers finished.")
}

func worker(ch chan int, wg *sync.WaitGroup, workerId int) {
	defer wg.Done()
	for {
		msg, ok := <-ch
		if !ok {
			fmt.Printf("Worker %d finished.\n", workerId)
			return
		}
		fmt.Printf("mes: %d; WID: %d\n", msg, workerId)
	}
}
