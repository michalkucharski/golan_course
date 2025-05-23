//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var color = "yellow"
	fmt.Println("My favorite color is", color)

	birthyear, age := 2000, 25
	fmt.Println("My birthyear is", birthyear, "my age is", age)

	var (
		first = "A"
		last  = "B"
	)

	fmt.Println("My first is: ", first)
	fmt.Println("And my last is", last)

	var ageInDays int

	ageInDays = 365 * age

	fmt.Println("My age in days:", ageInDays)

}
