package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3}
	var wg sync.WaitGroup
	wg.Add(len(nums))

	for _, num := range nums {
		go squares(num, &wg)
	}
	wg.Wait()
	fmt.Println("Main Routine..")
}

func squares(n int, wg *sync.WaitGroup) {
	fmt.Printf("Square of %d is: %d", n, n*n)
	fmt.Println()
	wg.Done()
}
