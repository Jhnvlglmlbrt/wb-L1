package main

import "fmt"

func removeElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func removeElementWOrder(slice []int, index int) []int {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func removeElementWithCopy(slice []int, index int) []int {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	indexToRemove := 2

	// С сохранением порядка элементов
	result1 := removeElement(numbers, indexToRemove)
	fmt.Println("With order:", result1)

	// Без сохранения порядка элементов
	result2 := removeElementWOrder(numbers, indexToRemove)
	fmt.Println("Without order:", result2)

	// С использованием copy (с сохранением порядка)
	result3 := removeElementWithCopy(numbers, indexToRemove)
	fmt.Println("With copy (order):", result3)

}
