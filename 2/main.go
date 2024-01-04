package main

import (
	"fmt"
	"sync"
	"time"
)

var arr = [5]int{2, 4, 6, 8, 10}

// для каждого числа запускаем горутину в которой выводим квадрат числа
// sleep для ожидания всех горутин
// передача значение через анонимную функцию в качестве аргумента в горутину для избежания гонки данных
func firstSolution() {
	for _, num := range arr {
		go func(number int) {
			fmt.Printf("%d | ", number*number)
		}(num)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println()
}

// использование небуферизированного канала
// для каждого числа из массива запускаем горутину, которая пишет в канал значение квадрата числа
// после чего читаем данные из канала
func secondSolution() {
	ch := make(chan int)
	for _, num := range arr {
		go func(number int) {
			ch <- number * number
		}(num)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d | ", <-ch)
	}
	fmt.Println()
}

// используем waitGroup и буферизированный канал
// каждого числа запускаем горутину, которая пишет в буфферизриованный канал квадрат числа
// ждем выполнения всех горутин с помощью wg.Wait(), закрываем канал и в цикле читаем все его значения
func thirdSolution() {
	ch := make(chan int, len(arr))
	var wg sync.WaitGroup
	for _, num := range arr {
		wg.Add(1)
		go func(number int) {
			defer wg.Done()
			ch <- number * number
		}(num)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for num := range ch {
		fmt.Printf("%d | ", num)
	}
}

func main() {
	firstSolution()
	secondSolution()
	thirdSolution()
}
