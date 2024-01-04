package main

import "fmt"

func main() {
	a, b := 10, 20
	fmt.Printf("before swap: a = %d, b = %d\n", a, b)

	a, b = b, a
	fmt.Printf("after sawp: a = %d, b = %d\n", a, b)
}
