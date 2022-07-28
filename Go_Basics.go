package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Declaration_styles() {
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

func hello() {
	fmt.Println("Hello World")
}

func print_and_format() {
	// printf is used for formatted printing

	// %T formatter is used for priting the type of the variable or a value
	fmt.Printf("Formatted type is: %T", 10)
	fmt.Println()
	// %v is used to print the actual value itself
	fmt.Printf("Formatted value is: %v", 10)
	fmt.Println()

	// %% is used for printing a percent literal
	fmt.Printf("Printing percent sign using formatter --> %%")
	fmt.Println()

	var bval bool = true
	fmt.Printf("Formatted boolean value is: %t", bval)

	// integer formatting
	// %b represents base2(binary) numbers..
	// %o erpresents base8(octal) numbers..
	// %d represents base10(decimal) numbers..
	// %x represents base16(hexadecimal) numbers..
	var x int16 = 1025
	fmt.Printf("Binary representation for 1025 is: %b", x)
	fmt.Println()
	fmt.Printf("Octal representation for 1025 is: %o", x)
	fmt.Println()
	fmt.Printf("Decimal representation for 1025 is: %d", x)
	fmt.Println()
	fmt.Printf("Hexadecimal representation for 1025 is: %x", x)
	fmt.Println()

	// float formatting
	// %e for scientific notation
	// %f / %F for decimal no exponent
	// %g for large exponents
	var float_num float32 = 53837485738.34298
	fmt.Printf("Scientific notation is?: %e", float_num)
	fmt.Println()
	fmt.Printf("Decimal with no exponent is: %f", float_num)
	fmt.Println()
	fmt.Printf("Large exponent formatting is: %g", float_num)
	fmt.Println()

	// string formatting
	var first_name string = "Ahsan"
	var last_name string = "Saif"
	fmt.Printf("First name is: %s and last name is: %s", first_name, last_name)
	fmt.Println()
	var d_quoted_string string = "Quoted_String"
	fmt.Printf("Formatting double quoted string: %q", d_quoted_string)
	fmt.Println()

	// width and padding
	// padding means margin from left
	// precision is the number if decimals to include in our floating point number
	// single . means just to round that number.. dont include the decimal part
	fmt.Printf("First name with padding is: %7s", first_name)
	fmt.Println()

	// floating precision..
	fmt.Printf("Padding and precision on floating number --> %10.3f", 428.347334)
	fmt.Println()
	fmt.Printf("Padding after %-9q %9q", first_name, last_name)
	fmt.Println()
	// padding with some digit
	fmt.Printf("Padding with 0: %05d", 45)
	fmt.Println()

	// concept of sprint..
	var out string = fmt.Sprintf("Padding with 0: %05d", 45)
	fmt.Println("Out is: ", out)
}

func variables_check() {
	var age uint = 23
	var name string = "Ahsan"
	fmt.Println(name)
	name = "Ahsan Saif"
	fmt.Println("Name is: ", name, "Age is: ", age)
}

func input() {
	Scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Type something: ")
	Scanner.Scan()
	// String input..
	inp := Scanner.Text()
	fmt.Println("You typed: ", inp)

	// converting string input to other types..
	int_inp, _ := strconv.ParseInt(Scanner.Text(), 10, 64)
	fmt.Printf("The number is: %d", int_inp)
}

func arithematic() {
	var num1 int = 4
	var num2 int = 8

	var x float32 = 9
	var y float32 = 2

	divi := x / y
	fmt.Printf("Result of division is: %v", divi)
	fmt.Println()

	// possible arithematic operators +, -, /, *, %, (, )
	// brackets are used for defining scope
	sum := num1 + num2
	fmt.Printf("Sum is: %v", sum)
}

func conditions_and_booleans() {
	var name1 string = "ahsan"
	var name2 string = "Ahsan"

	fmt.Printf("Sting less than check: %v", name1 > name2)
	fmt.Println()

	var x int = 6
	var y int = 7
	fmt.Printf("%v greater than %v is: %t", x, y, x > y)

}

func chained_confitions() {
	// or chained
	check := 6 < 7 || 7 > 10
	fmt.Printf("Or check is: %t", check)
	fmt.Println()

	// and check
	check1 := 6 < 7 && 7 > 10
	fmt.Printf("Or check is: %t", check1)
	fmt.Println()

	// not check
	check2 := !(6 < 7 && 7 > 10)
	fmt.Printf("Not check is: %t", check2)
	fmt.Println()

	fmt.Printf("Bool value is: %t", (true || false) && false)
	fmt.Println()
}

func conditionals() {
	age := 18
	if age > 18 {
		fmt.Printf("You are %v and eligible for license", age)
		fmt.Println()
	} else if age == 18 {
		fmt.Printf("You are %v just wait one more year", age)
		fmt.Println()
	} else {
		fmt.Printf("You are %v and below 18 hence not eligible for license", age)
		fmt.Println()
	}
}

func loops() {
	// for is written like below..
	for x := 0; x < 5; x++ {
		fmt.Println("Hi there")
	}
	y := 0
	// while is written as for in golang..
	for y < 1000 {
		if y%3 == 0 && y%6 == 0 {
			fmt.Println(y)
		}
		y++
	}
}

func switch_case() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter some number: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	switch input {
	case 1:
		fmt.Printf("Entered %v", input)
		fmt.Println()
	case 2:
		fmt.Printf("Entered %v", input)
		fmt.Println()
	default:
		fmt.Println("Entered something else")
	}

	ans := 7
	switch {
	case ans > 0:
		fmt.Println("Answer greater than 0")
	case ans < 0:
		fmt.Println("Answer less than 0")
	default:
		fmt.Println("Answer equals 0")
	}

}

func arrays() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Array is: %v", arr)
	fmt.Println()

	// partially assigned array..
	p_arr := [5]int{1, 2, 3}
	fmt.Printf("Partially assigned array is: %v", p_arr)
	fmt.Println()

	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Printf("Array sum is: %v", sum)
	fmt.Println()
}

func twoD_arrays() {
	arr_2D := [2][2]int{{1, 2}, {3, 4}}
	fmt.Printf("Tow dimentional array is: %v", arr_2D)
	fmt.Println()
}

func slices() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Array is: %v", arr)
	fmt.Println()

	var s []int = arr[1:3]
	fmt.Printf("Slice S is: %v", s)
	fmt.Println()
	fmt.Printf("Length of slice is: %v", len(s))
	fmt.Println()
	fmt.Printf("Capacity of slice is: %v", cap(s))
	fmt.Println()
	fmt.Println("Entending slice..")
	// slice originally made of arr extending it to its capacity..
	s = s[:cap(s)]
	fmt.Printf("Extended slice is: %v", s)
	fmt.Println()

	// Declaring and initializing a slice..
	var sl []int = []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("Slice is: %v", sl)
	fmt.Println()

	// append to a slice (return type.. returns a new slice)
	b := append(sl, 7)
	fmt.Printf("Slice b is: %v", b)
	fmt.Println()

	// making an array using make..
	a := make([]int, 5)
	fmt.Printf("Array after make is: %v", a)
	fmt.Println()
}

func slice_array_range() {
	var a []int = []int{1, 3, 5, 7, 9, 11, 13, 15}
	for i, element := range a {
		fmt.Printf("Index is: %d and element is: %d", i, element)
		fmt.Println()
	}

	// concept of underscore while using range keyword..
	// underscore is an anonymous variable..
	// value of underscore can't be accessed as it is anonymous..
	// printing just elements using range keyword..
	for _, element := range a {
		fmt.Printf("Element: %d\n", element)
	}

	// printing just indexes using range keyword..
	for i := range a {
		fmt.Printf("Index: %d\n", i)
	}
}

func maps() {
	// declaring maps
	var mp map[string]int = map[string]int{
		"Apples":  1,
		"Oranges": 2,
		"Pear":    3,
	}
	// printing map..
	fmt.Print("Map is: ", mp, "\n")
	// inserting/change
	mp["Banana"] = 4
	mp["Apples"] = -1
	fmt.Print("Now the map is: ", mp, "\n")
	fmt.Println()

	// other way of declaring map..
	new_mp := make(map[string]int)
	new_mp["Ahsan"] = 1
	fmt.Println("New map is: ", new_mp)

	// deleting from a map..
	delete(mp, "Apples")
	fmt.Print("Map after deleting apples is: ", mp, "\n")

	// checking if the key exists in a map..
	val, ok := mp["Apples"]
	fmt.Printf("Value is: %d and ok is: %v", val, ok)
}

func add(x, y int) int {
	return x + y
}

// return squares of 2 numbers..
func squares(x, y int) (z1 int, z2 int) {
	// z1 and z2 are labelled return values
	defer fmt.Println("Defer happs at the end of function..")
	z1 = x * x
	z2 = y * y
	fmt.Println("Before hello..")
	return z1, z2
}

func referencing() {
	fmt.Println("Hellowww")
}

func func_calling_function(func1 func(int) int) {
	fmt.Println("Square value is: ", func1(5))
}

func returnFunc(x string) func() {
	return func() { fmt.Println(x) }
}

func jsonConversion() {

}

func main() {
	// hello()
	// input()
	// arithematic()
	// conditions_and_booleans()
	// chained_confitions()
	// conditionals()
	// loops()
	// switch_case()
	// arrays()
	// twoD_arrays()
	// slices()
	// slice_array_range()
	// maps()
	// sum := add(5, 9)
	// fmt.Printf("Sum is: %d\n", sum)
	// x_sqr, y_sqr := squares(5, 9)
	// fmt.Printf("5 square is: %d and 9 square is: %d\n", x_sqr, y_sqr)

	// assigning function to a variable..
	func_ref := referencing
	func_ref()

	// function inside function..
	test := func() {
		fmt.Print("Function inside function..\n")
	}
	test()

	// direct function call after declaring it
	negation := func(x int) int {
		return x * (-1)
	}(8)
	fmt.Printf("Negation function called right after declaration: %d\n", negation)

	square := func(x int) int {
		return x * x
	}
	func_calling_function(square)
	returnFunc("Ahsan")()
}
