package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count int
	lock  sync.Mutex
)

func main() {
	iterations := 10000000
	for i := 0; i < iterations; i++ {
		go iterate()
	}
	time.Sleep(time.Second)
	fmt.Println("Count is: ", count)
}

func iterate() {
	lock.Lock()
	count++
	lock.Unlock()
}
