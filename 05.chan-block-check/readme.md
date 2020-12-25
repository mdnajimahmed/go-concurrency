```
func main() {
	ch := make(chan int)
	go func(myChan chan<- int) {
		
	}(ch)

	out := <-ch
	fmt.Println("Value received!", out) 
}
```
- Line is bocking, waiting *INSIDE MAIN GO ROUTINE* for some message to arrive. But all other go routines finished execution, hence we deterministically know that this waiting is forever, nothing is gonna happen! That's why its deadlock. 

```
func main() {
	ch := make(chan int)
	go func(myChan chan<- int) {
		close(myChan)
	}(ch)

	out := <-ch
	fmt.Println("Value received!", out)
}
```
- We did not send any message to channel, we just closed it. When we close a channel a default message (default value of channel's type) is sent to channel which is recevied by the main go routine, therefore the bocking/waiting state ends at main go routine and the program terminates without panic.

```
func main() {
	ch := make(chan int)
	go func(myChan chan<- int) {
		myChan <- 1
		myChan <- 2
		myChan <- 3
	}(ch)

	out := <-ch
	fmt.Println("Value received!", out) 
}
```

- We can put more than one value which is going to be ignored, main routine's `out := <-ch` instruction waits until it receives one msg from the channel and as soon as it receives one message it continues exectuion, then main routine terminitates so the program terminates(may be from a un desired state,e.g a go routine was forefully shut down before finishing its job)