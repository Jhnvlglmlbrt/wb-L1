package main

import (
	"fmt"
)

func checkType(val interface{}) {
	switch val.(type) {
	case int:
		fmt.Println("Type: int")
	case string:
		fmt.Println("Type: string")
	case bool:
		fmt.Println("Type: bool")
	case chan int:
		fmt.Println("Type: channel of int")
	default:
		fmt.Println("Type: unknown")
	}
}

func main() {
	var a interface{} = 42
	var b interface{} = "Hello"
	var c interface{} = true
	var d interface{} = make(chan int)

	checkType(a)
	checkType(b)
	checkType(c)
	checkType(d)
}
