package main

import (
	"fmt"
	"time"
)

func procudeOdd(ch chan<- int) {
	ch <- 1
	time.Sleep(1 * time.Second)
	ch <- 3
	time.Sleep(1 * time.Second)
	ch <- 5
	time.Sleep(1 * time.Second)
}
func procudeEven(ch chan<- int) {
	ch <- 2
	time.Sleep(2 * time.Second)
	ch <- 4
	time.Sleep(2 * time.Second)
	ch <- 6
	time.Sleep(2 * time.Second)
	//close(ch)
}
func main() {
	oddCh := make(chan int)
	evenCh := make(chan int)
	go procudeOdd(oddCh)
	go procudeEven(evenCh)
	// select {
	// case odd := <-oddCh:
	// 	{
	// 		fmt.Println(odd)
	// 	}
	// case even := <-evenCh:
	// 	{
	// 		fmt.Println(even)
	// 	}
	// case <-time.After(5 * time.Second):
	// 	{
	// 		fmt.Println("I cant wait any more,sorry!")
	// 	}
	// 	// default:
	// 	// 	fmt.Println("I am default")
	// }
	/*** Selects ends as soon as one message is arrived*/
	/* dont close channel which has select on it!, smells bad*/
loop:
	for {
		select {
		case odd := <-oddCh:
			{
				fmt.Println(odd)

			}
		case even := <-evenCh:
			{
				fmt.Println(even)
			}
		case <-time.After(5 * time.Second):
			{
				fmt.Println("I cant wait any more,sorry!")
				break loop
			}
			// default:
			// 	fmt.Println("I am default")
		}
	}

}
