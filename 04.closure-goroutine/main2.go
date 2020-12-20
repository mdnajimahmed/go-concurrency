package main

import (
	"fmt"
	"sync"
)

// testing w8 group
func main() {
	var w sync.WaitGroup

	w.Add(1)
	w.Add(1)

	go func() {
		fmt.Println("I am done once")
		w.Done()
	}()
	/*** Calling add twice on w is equivalent to w.Add(2)*/
	/* Keeping this commented out program will not run*/

	go func() {
		fmt.Println("I am done once")
		w.Done()
	}()

	w.Wait()
}
