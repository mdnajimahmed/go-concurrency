package main

import (
	"fmt"
)

func raisePanicLev3() {
	panic("I panic for no reason")
	// panic(nil)
}

func raisePanicLev2() {
	fmt.Println("before raisePanicLev3")
	raisePanicLev3()
	fmt.Println("after raisePanicLev3") // will not get called
}

func raisePanicLev1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f") // panic(nil) nill not print this log!!! :p
		}
	}()
	fmt.Println("before raisePanicLev2")
	raisePanicLev2()
	fmt.Println("after raisePanicLev2") // tricky one - will also not get called
}

func main() {

	/* Panic immidiately stops execution regardless of go routine*/

	fmt.Println("before raisePanicLev1")
	raisePanicLev1()
	fmt.Println("after raisePanicLev1") // will get called
}

// // after execution, before return -> defer
/*** PANIC HAPPENS AFTER DEFERRED EXECUTIONS ARE HAPPENED

Here is the execution order
- Exection of the function , (closing curly brace)
- execute defer
- panic
- return value


recover will receive nil if
- No panic happened
- Panic happened with nil argument
- Panic happened outside of current go routine
***/
