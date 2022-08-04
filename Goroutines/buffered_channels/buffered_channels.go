package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Buffered channels")
	defer fmt.Println("Total time is: ", time.Since(start))

	channel := make(chan string, 2)
	channel <- "First Message"
	channel <- "Second Message"
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
