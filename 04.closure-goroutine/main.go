package main

import (
	"fmt"
	"sync"
	"time"
)

var w sync.WaitGroup

// not safe, because it undetermnisticaly picks up the value of num
func testGoClousure() {
	num := 0

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Num b4", num)
		num++
		fmt.Println("Num aft", num)
		w.Done()
	}()
	num = 1
	fmt.Println("Finished execution! num = ", num)
}

func testGoClousureScoped() {
	num := 0

	go func(num int) { // not a good variable name, use different name e.g nameCopy
		time.Sleep(5 * time.Second)
		fmt.Println("Num b4", num)
		num++
		fmt.Println("Num aft", num)
		w.Done()
	}(num)
	num = 1
	fmt.Println("Finished execution! num = ", num)
}

// func main() {
// 	w.Add(1)
// 	// testGoClousure()
// 	testGoClousureScoped()
// 	w.Wait()
// }
