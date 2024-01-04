package main

import (
	"fmt"
	"strings"
)

var justString string

// переслайсинг (re-slicing) среза не создаёт копию массива в основании. Массив полностью будет хранимться в памяти,
// пока на него не перестанут ссылаться. Иногда это вызывает хранение всех данных в памяти,когда нужна только их небольшая часть.
// Так как слайс ссылается на исходный массив, пока слайс есть в памяти, сборщик мусора не сможет очистить массив

func createHugeString(s int) string {
	return strings.Repeat("s", s)
}

func someFunc() {
	// большая строка будет в переменной
	v := createHugeString(1 << 10)

	// слайс вместимостью 100 байт
	slice := make([]byte, 100)

	// копируем фрагмент огромной строки в заранее подготовленый слайс
	copy(slice, v[:100])
	justString = string(slice)
}

func main() {
	someFunc()
	fmt.Println(justString)
}
