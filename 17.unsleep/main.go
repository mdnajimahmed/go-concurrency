package main

// import (
// 	"fmt"
// 	"time"
// )

// func routineToSleep(ch <-chan bool) {
// 	fmt.Println("Begin of routineToSleep")
// 	now := time.Now()
// 	time.Sleep(5 * time.Second)
// 	// if <-ch {
// 	// 	fmt.Println("Hey I woke up forcefully !")
// 	// } else {
// 	// 	fmt.Println("Hey I woke up normally!")
// 	// }
// 	select {
// 	case <-ch:
// 		fmt.Printf("Hey I woke up forcefully after %v !", time.Now().Sub(now))
// 	default:
// 		fmt.Println("I am default")
// 	}
// 	fmt.Println("End of routineToSleep")
// }
// func main() {
// 	fmt.Println("Hi")
// 	ch := make(chan bool)
// 	fmt.Println("Begin of main")
// 	go routineToSleep(ch)
// 	select {
// 	case <-time.After(2 * time.Second):
// 		ch <- true
// 	}
// 	fmt.Println("End of main")
// }
