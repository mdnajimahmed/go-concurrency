package main

import (
	"fmt"
	"time"
)

func blockingVersion() {
	ch := make(chan int)
	fmt.Println("Before inserting into channel")
	ch <- 5
	fmt.Println("Before inserting into channel")

	fmt.Println("Before polling from channel")
	out := <-ch
	fmt.Println("After polling from channel. value = ", out)

}

func nonBlockingChannel() {

	/*
		Buffered channel with buffer size = 1.
		This channel will not block until it receives 1 item.
		It it receives more than one item, it will block that go routine,
		and hope another go routine will pick up items.
	*/
	ch := make(chan int, 1)

	fmt.Println("Before inserting into channel")
	ch <- 5 // non blocking,this message will be buffered to be picked up later
	fmt.Println("Before inserting into channel")

	fmt.Println("Before polling from channel")
	out := <-ch
	fmt.Println("After polling from channel. value = ", out)
}

func example() {
	producer := func() chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 0; i < 5; i++ {
				ch <- (i + 100)
				time.Sleep(3 * time.Second)
			}
		}()
		fmt.Println("Confusing the audience")
		return ch
	}

	source := producer()

	for v := range source {
		fmt.Println(v)
	}
}

func main() {
	// blockingVersion()
	// nonBlockingChannel()
	example()
}
