//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.

func greeting(name string) {
	fmt.Println("Hello ", name)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func someMsg() {
	fmt.Println("Here is some message !!")
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer

func threeArgsSum(arg1, arg2, arg3 int) int {
	return arg1 + arg2 + arg3
}

//* Write a function that returns any number

func oneNumber() int {
	return 4
}

// * Write a function that returns any two numbers
func twoNumbers() (int, int) {
	return 45, 56
}

func main() {

	//* Add three numbers together using any combination of the existing functions.
	//  * Print the result

	sumthree := threeArgsSum(oneNumber(), 4, 6)

	fmt.Println("SUm of three numbers:", sumthree)

	//* Call every function at least once
	greeting("Lola")

	someMsg()

}
