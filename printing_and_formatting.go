package main

import "fmt"

// working on sprintf anf printf
func main() {
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
