package main

import (
	"fmt"
	"sync"

	/*VVI use module name / package name!!! I missed this!*/
	"timelyship.com/learning/sync-once/domain"
	"timelyship.com/learning/sync-once/domain2"
)

func main() {
	fmt.Println("Hi!")
	var w sync.WaitGroup
	w.Add(3)
	for i := 0; i < 3; i++ {
		go func(itr int) {
			defer w.Done()
			person1 := domain.NewPerson()
			fmt.Printf("%d Person-1 %p\n", itr, person1)
			person2 := domain2.NewPerson()
			fmt.Printf("%d Person-2 %p\n", itr, person2)
		}(i)
	}
	w.Wait()
}
