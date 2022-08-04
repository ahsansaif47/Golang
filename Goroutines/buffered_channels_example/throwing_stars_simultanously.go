package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	defer fmt.Println("Total time is: ", time.Since(start))

	sChannel := make(chan string)
	go yellScores(sChannel)

	// customized for loop breaks when the channel closes..
	for {
		message, open := <-sChannel
		if !open {
			break
		}
		fmt.Println(message)
	}

}

func yellScores(starChannel chan string) {
	// Updating channel info with scores..
	rand.Seed(time.Now().UnixNano())
	ninjaStars := 3
	for i := 0; i < ninjaStars; i++ {
		score := rand.Intn(10)
		starChannel <- fmt.Sprint("You scored: ", score)
	}
	// closing channel at the end..
	// after sending all the messages..
	close(starChannel)
}
