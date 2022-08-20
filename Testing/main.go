package main

import "fmt"

func addTwo(x int) int {
	return x + 2
}

func main() {
	fmt.Println("Go Testing Tutorial")
	result := addTwo(2)
	fmt.Println("Adding 2 to 2 resulted in: ", result)
}
