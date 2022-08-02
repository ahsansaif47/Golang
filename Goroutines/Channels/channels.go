package main

import (
	"fmt"
	"time"
)

func main() {
	println("Channels Intro..")
	// Keeping track of time..
	start := time.Now()
	evilNinja := "Tim"
	defer fmt.Println("Time taken by golang ninja for the mission is: ", time.Since(start))

	// Declaring and initializing a channel named attackSignal..
	attackSignal := make(chan bool)
	go attack(evilNinja, attackSignal)

	// channel can be reused by reassigning the value..
	// if the value is updated true. we can reuse the channel by making it false..
	// likt attackSignal <- false
	fmt.Println("Signal: ", <-attackSignal)
}

func attack(target string, signal chan bool) {
	time.Sleep(time.Second)
	fmt.Println("Attacking ", target)
	signal <- true
}
