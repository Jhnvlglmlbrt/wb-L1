package main

import "fmt"

func quickSort(list []int) []int {
	if len(list) < 2 {
		return list
	} else {
		pivot := list[0]

		less := []int{}
		greater := []int{}

		for _, num := range list[1:] {
			if pivot > num {
				less = append(less, num)
			} else {
				greater = append(greater, num)
			}
		}

		less = append(quickSort(less), pivot)
		greater = quickSort(greater)
		return append(less, greater...)
	}
}

func main() {
	fmt.Println(quickSort([]int{10, 5, 2, 3}))
}
