package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Concurrency")
	start := time.Now()
	evilNinjas := []string{"Bob", "Tim", "Alice"}
	defer fmt.Println("Time taken is: ", time.Since(start))

	// attacking ninjas..
	for _, n := range evilNinjas {
		go attack(n)
	}

	/*
		For preventing the program from exiting from
		executing the threads you need to add more than
		a second delay before the main exits..
	*/
	time.Sleep(time.Second * 2)
}

func attack(target string) {
	fmt.Println("Throwing ninja stars at: ", target)
	time.Sleep(time.Second)
}
