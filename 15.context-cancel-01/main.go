package main

import (
	"context"
	"fmt"
)

func generate(ctx context.Context) <-chan int {
	out := make(chan int)
	n := 1
	go func() {
		defer close(out)
		for {
			select {
			case out <- n:
			case <-ctx.Done():
				return
			}
			n++
		}
	}()
	return out

}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	for number := range generate(ctx) {
		fmt.Println(number)
		if number == 5 {
			cancel()
		}
	}
}
