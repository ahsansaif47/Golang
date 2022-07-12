package main

import "fmt"

func main() {
	fmt.Println("Go program")
	// implicit variable declaration and assignment
	// golang guesses the type if the variable
	var x = 260
	// checking the type of variable
	fmt.Printf("Type of variable x is: %T", x)
	fmt.Println()

	// explicit varibale declaration and assignment
	var y int8 = 24
	fmt.Printf("Type of variable y is: %T", y)
	fmt.Println()

	// walrus operator (omits the usage of word var while declaring a variable)
	number := 24
	fmt.Print("Number is: ", number)

	// checking default values for datatypes..
	var v1 uint8
	fmt.Println("Default value for v1 variable is: ", v1)
	var v2 float32
	fmt.Println("Default value for v2 variable is: ", v2)
	var v3 bool
	fmt.Println("Default value for v3 variable is: ", v3)
	var v4 string
	fmt.Println("Default value for v4 variable is: ", v4)
	// updating the dafault value for string datatype
	v4 = "Ahsan"
	fmt.Println("Initialize value for v4 variable is: ", v4)

}
