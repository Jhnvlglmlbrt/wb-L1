package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

type AtomicCounter struct {
	counter int64
}

func (ac *AtomicCounter) Increment() {
	atomic.AddInt64(&ac.counter, 1)
}

func main() {
	// используем структуру с мьютексом в которой хранится наш счетчик
	// при каждой операции с числом блокируем мьютекс

	var wg sync.WaitGroup

	c := &Counter{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				c.Increment()
			}
		}()
	}

	wg.Wait()

	fmt.Println(c.value)

	// использование атомарной переменной вместо мьютексов
	ac := &AtomicCounter{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				ac.Increment()
			}
		}()
	}
	wg.Wait()

	fmt.Println(ac.counter)
}
