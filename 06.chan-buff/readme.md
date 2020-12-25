```
ch := make(chan int)
	fmt.Println("Before inserting into channel")
	ch <- 5
	fmt.Println("Before inserting into channel")

	fmt.Println("Before polling from channel")
	out := <-ch
	fmt.Println("After polling from channel. value = ", out)
```

- Outputs this
```
go run main.go
Before inserting into channel
fatal error: all goroutines are asleep - deadlock!

```

Because main routine is waiting some other routine to pick up the value which is just inserted in channel, but the go runtime detected that there is no hope! The only go routine(main) is blocked! Thats why it panics.
