package main

import (
	"fmt"
	"strings"
)

// START OMIT
func splitAddress(address string) (string, string) {
	parts := strings.Split(address, ":")
	return parts[0], parts[1]
}

func main() {
	host, port := splitAddress("localhost:8080")
	fmt.Println(host)
	fmt.Println(port)
}

// END OMIT
