/*
	Objective:
		To convert JSON string to json onject.
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Ninja struct {
	Name   string
	Weapon string
	Level  int
}

func main() {
	fmt.Println("Working with JSON")
	jString := `{"name": "Wallace", "Weapon":"Sword", "level":1}`
	fmt.Printf("JSON string is: %v", jString)
	fmt.Println()

	// Creating fo object of type Ninja
	var wallace Ninja
	// JSON string to ho object..
	err := json.Unmarshal([]byte(jString), &wallace)

	if err == nil {
		fmt.Println("JSON Object is: ", wallace)
	} else {
		fmt.Println("Encountered with ERROR ", err)
	}

	// Go object to josn string..
	jStr, err := json.Marshal(wallace)
	// jStr is a byte slice..
	// It can be formatted as string using %s

	fmt.Printf("JSON String is: %s", jStr)
	fmt.Println()
}
