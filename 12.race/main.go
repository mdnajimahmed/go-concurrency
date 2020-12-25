package main

import (
	"fmt"
	"math/rand"
	"time"
)

func problemCode() {
	fmt.Println("Hi")
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(randomDuration())
	// }
	start := time.Now()
	fmt.Println("Start", start)
	var t *time.Timer
	/*time after function returns a value, and exectues the anonymous function in a separate
	go routine, these two operations are asynchronous, hence its possibe that the anonymous func
	is running before t = time.AfterFunc actually finished assigning the value to t! So? inside
	go routine, which is running async can fail in line t.Reset(randomDuration()) because t might be
	nil!
	*/
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println("time elapsed after start: ", time.Now().Sub(start))
		t.Reset(randomDuration())
	})
	time.Sleep(5 * time.Second)
}

func solutionCode() {
	fmt.Println("Hi")
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(randomDuration())
	// }
	ch := make(chan bool)
	start := time.Now()
	fmt.Println("Start", start)
	var t *time.Timer

	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println("time elapsed after start: ", time.Now().Sub(start))
		ch <- true

	})

	for time.Since(start) < 5*time.Second {
		<-ch
		/*Assignment and reset is in the same go routine that means they are automaticaly synchronized*/
		t.Reset(randomDuration())
	}
}
func main() {
	// problemCode()
	solutionCode()

}

func randomDuration() time.Duration {
	f := 2 * rand.Float32() * float32(time.Second)
	return time.Duration(f)
}
