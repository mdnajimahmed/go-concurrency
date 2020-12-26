package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// producing gold is an expensive operation
var w sync.WaitGroup

func mineGold(ctx context.Context, gold chan<- int) {
	now := time.Now()
	defer close(gold)
	select {
	case <-time.After(5 * time.Second):
		gold <- 1000000
	case <-ctx.Done():
		fmt.Printf("Hey I woke up forcefully after %v!\n", time.Now().Sub(now))
	}
	w.Done()
}

func main() {
	now := time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("Begin of main")
	gold := make(chan int)
	w.Add(1)
	go mineGold(ctx, gold)
	select {
	case g := <-gold:
		fmt.Println("Mined gold of amount ", g)
	case <-time.After(2 * time.Second):
		fmt.Printf("Going to cancel from main after %v !\n", time.Now().Sub(now))
		cancel()
	}
	fmt.Println("End of main")
	w.Wait()
	fmt.Println("NumGoroutine", runtime.NumGoroutine())
}
