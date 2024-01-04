package main

import "fmt"

func binarySearch(list []int, i int) int {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		if list[mid] == i {
			return mid
		} else if list[mid] < i {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 3))
}
