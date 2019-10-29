package main

import (
	"fmt"
	"time"
)

func squares(c chan int) {
	time.Sleep(3 * time.Second)
	c <- 5
}

func main() {
	fmt.Println("Program started")

	c := make(chan int)

	go squares(c)

	select {
	case m := <-c:
		fmt.Println(m)
	case <-time.After(2 * time.Second):
		fmt.Println("timed out")
	}

	fmt.Println("Program finished")
}
