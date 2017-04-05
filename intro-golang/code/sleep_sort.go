package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// START OMIT
	var wg sync.WaitGroup
	x := []int{3, 1, 4, 2, 5}

	for _, i := range x {
		go func(a int) {
			wg.Add(1)
			defer wg.Done()
			time.Sleep(time.Duration(a) * time.Second)
			fmt.Println(a)
		}(i)
	}
	wg.Wait()
	// END OMIT
}
