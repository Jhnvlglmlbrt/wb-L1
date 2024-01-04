package main

import (
	"fmt"
	"time"
)

func mySleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Start")
	mySleep(3)
	fmt.Println("End after 3 seconds")
}
