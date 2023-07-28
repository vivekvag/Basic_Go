package main

import (
	"fmt"
	"time"
)

func main() {
	one := make(chan string)
	two := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		one <- "Hey"
	}()

	go func() {
		time.Sleep(4 * time.Second)
		two <- "Hello"
	}()

	select {
	case rec1 := <-one:
		fmt.Println("I received from channel one!", rec1)
	case rec2 := <-two:
		fmt.Println("I received from channel two!", rec2)
	}

}
