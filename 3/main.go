package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// операция увеличения выполняется как одна неделимая (атомарная) операция и гарантирует потокобезопасность
// из-за гонки данных, ведь много горутин запускается в цикле
// чтобы избежать конфликта между горутинами, передаём значение num внутрь замыкания как аргумент
func main() {
	var res uint64
	var wg sync.WaitGroup

	arr := []uint64{2, 4, 6, 8, 10}
	for _, num := range arr {
		wg.Add(1)
		go func(n uint64) {
			defer wg.Done()
			atomic.AddUint64(&res, n*n)
		}(num)
	}

	wg.Wait()
	fmt.Println(res)
}
