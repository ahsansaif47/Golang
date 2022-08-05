/*
	Learning Objective:
		To get to know about advanced functions
		Like passing a function to another function as a parameter
*/

package main

import "fmt"

func returnVal(i int) int {
	return i
}

// Passed function as a parameter..
func returnFunction(i int, f func(int) int) int {
	return f(i)
}

func not_main() {
	fmt.Println("Advaced functions")
	// Printing function..
	fmt.Printf("%T", returnVal)
	fmt.Println()

	// Calling function..
	fmt.Printf("%T", returnVal(5))
	fmt.Println()

	// Concept of function returning other function..
	fmt.Printf("Return type is: %T", returnFunction(5, returnVal))
	fmt.Println()
	fmt.Printf("Value returned is: %v", returnFunction(5, returnVal))
	fmt.Println()

	// Anonymous functions declared here..
	anonVar := func() {
		fmt.Println("Anomymous function")
	}
	fmt.Println("Hello World")
	// Acts on being called..
	anonVar()
}
