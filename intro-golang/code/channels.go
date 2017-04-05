package main

import (
	"fmt"
)

func main() {
	// START OMIT
	ch := make(chan string)

	go func() {
		ch <- "Hello World!"
	}()

	val := <-ch
	fmt.Printf("Received: %v\n", val)
	// END OMIT
}
