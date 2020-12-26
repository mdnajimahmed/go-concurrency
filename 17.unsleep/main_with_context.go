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

func mineGold(ctx context.Context) <-chan int {
	now := time.Now()
	out := make(chan int)
	go func() {
		defer close(out)
		select {
		case <-time.After(5 * time.Second):
			out <- 1000000
		case <-ctx.Done():
			fmt.Printf("Hey I woke up forcefully after %v!\n", time.Now().Sub(now))
		}
		w.Done()
	}()

	return out
}

func main() {
	now := time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("Begin of main")
	w.Add(1)
	select {
	case g := <-mineGold(ctx):
		fmt.Println("Mined gold of amount ", g)
	case <-time.After(2 * time.Second):
		fmt.Printf("Going to cancel from main after %v !\n", time.Now().Sub(now))
		cancel()
	}
	w.Wait()
	fmt.Println("NumGoroutine", runtime.NumGoroutine())
	fmt.Println("End of main")

}
