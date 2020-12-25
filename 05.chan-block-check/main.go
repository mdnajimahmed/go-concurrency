package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func(myChan chan<- int) {
		myChan <- 1
		myChan <- 2
		myChan <- 3

		// time.Sleep(5 * time.Second)
		// I do nothing
		// close(myChan)
	}(ch)

	out := <-ch
	fmt.Println("Value received!", out) //should be deadlock
}
