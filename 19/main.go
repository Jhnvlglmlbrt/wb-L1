package main

import (
	"fmt"
	"strings"
)

func reverseString(str string) string {
	// руны позволяют корректно обрабатывать многобайтные Unicode-символы
	runes := []rune(str)
	var b strings.Builder

	for i := len(runes) - 1; i >= 0; i-- {
		b.WriteRune(runes[i])
	}

	return b.String()
}

func main() {
	fmt.Println(reverseString("string !+- test 日本語"))
}
