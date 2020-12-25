package main

import (
	"fmt"
	"sync"
	"time"
)

var w sync.WaitGroup
var sharedRsc = make(map[string]string)

func consumer1(c *sync.Cond) {
	defer w.Done()
	c.L.Lock()
	for len(sharedRsc) == 0 {
		fmt.Println("I need to wait(consumer1)")
		c.Wait()
	}
	fmt.Println("consumer1 : ", sharedRsc["rsc"])
	c.L.Unlock()
}

func consumer2(c *sync.Cond) {
	defer w.Done()
	c.L.Lock()
	for len(sharedRsc) == 0 {
		fmt.Println("I need to wait(consumer2)")
		c.Wait()
	}
	fmt.Println("consumer2 : ", sharedRsc["rsc"])
	c.L.Unlock()
}

func producer(c *sync.Cond) {
	defer w.Done()
	c.L.Lock()
	fmt.Println("Producer taking a 3 sec pause")
	time.Sleep(3 * time.Second)
	sharedRsc["rsc"] = "foo"
	// c.Signal()
	c.Broadcast()
	c.L.Unlock()
}

func main() {
	mu := sync.Mutex{}
	c := sync.NewCond(&mu)
	w.Add(3)
	go consumer1(c)
	go consumer2(c)
	go producer(c)
	w.Wait()

}
