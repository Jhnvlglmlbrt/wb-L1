package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// 1. С использованием контекста
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Context done. Goroutine finished.")
		default:
			fmt.Println("Goroutine working.")
		}
	}()

	// Отмена контекста через 2 секунды
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond) // чтобы горутина успела завершиться

	// 2. С использованием канала
	stopCh := make(chan struct{})

	go func() {
		select {
		case <-stopCh:
			fmt.Println("Stop signal received. Goroutine finished.")
		default:
			fmt.Println("Goroutine working.")
		}
	}()

	time.Sleep(2 * time.Second)
	close(stopCh)
	time.Sleep(500 * time.Millisecond)

	// 3. С использованием WaitGroup
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine working.")
	}()

	wg.Wait()

	// 4. С использованием таймера
	go func() {
		timer := time.NewTimer(2 * time.Second)
		select {
		case <-timer.C:
			fmt.Println("Timer expired. Goroutine finished.")
		default:
			fmt.Println("Goroutine working.")
		}
	}()

	time.Sleep(500 * time.Millisecond)
}
