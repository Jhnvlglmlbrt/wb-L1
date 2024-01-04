package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	arr := strings.Split(str, " ")

	var b strings.Builder

	for i := len(arr) - 1; i >= 0; i-- {
		b.WriteString(" ")
		b.WriteString(arr[i])
	}

	return b.String()
}

func main() {
	str := "snow dog sun"
	fmt.Printf("    init string: %s\n", str)
	fmt.Printf("reversed string: %s\n", strings.Trim(reverseWords("snow dog sun"), " "))
}
