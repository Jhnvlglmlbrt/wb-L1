package main

import "fmt"

func main() {
	arr := [5]int{1, 5, 3, 4, 2}
	arr1 := [6]int{4, 7, 12, 5, 10, 8}

	same := map[int]bool{}

	for _, num := range arr1 {
		same[num] = true
	}

	for _, num := range arr {
		if same[num] {
			fmt.Println("Same number: ", num)
		}
	}
}
