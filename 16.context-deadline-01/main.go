package main

import (
	"context"
	"fmt"
	"time"
)

func compute(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		deadline, ok := ctx.Deadline()
		if ok {
			fmt.Println("Deadline set to " + deadline.String())
			// 02. task esitimated that it needs  2500 ms / 2.5 sec
			expEndTimeOfComputation := time.Now().Add(2500 * time.Millisecond)
			if deadline.Sub(expEndTimeOfComputation) < 0 {
				fmt.Println("Not sufficient time given ... terminating")
				return
			}
			fmt.Println("I have enough time, I am starting processing," +
				"but I could be terminated while I am doing processing")
		}
		fmt.Println("Task started")
		time.Sleep(2500 * time.Millisecond)
		fmt.Println("Task finished")
		select {
		case out <- 100: // 100 is the result(fake/dummy) of the task
		case <-ctx.Done():
			return
		}
	}()
	return out

}
func main() {
	// 01. task must be finihsed within 3000 ms / 3 sec
	deadline := time.Now().Add(3000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	select {
	case val := <-compute(ctx):
		fmt.Println("Processing finished! vaue = ", val)
	/* Cancel after 1500 ms. To stop cancellation comment out this case*/
	case <-time.After(1500 * time.Millisecond):
		fmt.Println("Cancelling after 1500 Millisecond")
		// fmt.Println("vaue = ", val)
		cancel()
	}

}
