package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// START OMIT
	var wg sync.WaitGroup
	inChan := make(chan int)

	// Consumers
	for i := 0; i < 3; i++ {
		go func(id int) {
			wg.Add(1)
			defer wg.Done()
			for job := range inChan {
				time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
				fmt.Printf("Consumer with id %v, done with job: %v\n", id, job)
			}
		}(i)
	}

	// Producers
	for i := 0; i < 20; i++ {
		inChan <- i
	}
	close(inChan)

	wg.Wait()
	// END OMIT
}
