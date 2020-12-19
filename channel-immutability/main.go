package main

import (
	"fmt"
	"sync"
	"time"
)

var w sync.WaitGroup

// Book is a dummy type to test immutability
type Book struct {
	page int
}

func main() {
	ch := make(chan Book)
	w.Add(2)
	go produce(ch)
	go consume(ch)
	w.Wait()
}

func consume(ch <-chan Book) {
	b := <-ch
	fmt.Printf("Address in consume %p\n", &b) // Address in consume 0xc00030a040 , SEE!!! its different.
	w.Done()
}

func produce(ch chan<- Book) {
	b := Book{
		page: 10,
	}
	fmt.Printf("Address in produce %p\n", &b) // Address in produce 0xc000474740
	time.Sleep(1 * time.Second)
	ch <- b
	w.Done()
}
