package main

import (
	"fmt"
)

func main() {
	var num int64

	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Ошибка ввода числа:", err)
		return
	}

	var bitValue int
	fmt.Print("Введите значение бита (0 или 1): ")
	_, err = fmt.Scan(&bitValue)
	if err != nil || (bitValue != 0 && bitValue != 1) {
		fmt.Println("Ошибка ввода значения бита.")
		return
	}

	var bitPosition int
	fmt.Print("Введите индекс бита: ")
	_, err = fmt.Scan(&bitPosition)
	if err != nil {
		fmt.Println("Ошибка ввода индекса бита:", err)
		return
	}

	num = setBit(num, uint(bitPosition), bitValue)

	fmt.Printf("Число после установки бита %d в %d: %d\n", bitPosition, bitValue, num)
}

// setBit устанавливает i-й бит в указанное значение (0 или 1)
func setBit(num int64, pos uint, value int) int64 {
	// Очистить i-й бит, используя битовую маску
	num &^= int64(1) << pos

	// Установить i-й бит в указанное значение
	num |= int64(value) << pos

	return num
}
