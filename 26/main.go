package main

import (
	"fmt"
	"strings"
)

func areAllCharactersUnique(s string) bool {
	seen := make(map[rune]bool)
	for _, char := range strings.ToLower(s) {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

func main() {
	stringsToCheck := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, s := range stringsToCheck {
		result := areAllCharactersUnique(s)
		fmt.Printf("str: %s, unique: %t\n", s, result)
	}
}
