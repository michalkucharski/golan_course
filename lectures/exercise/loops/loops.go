//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {

	number := 1

	for number < 50 {
		if number%3 == 0 && number%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if number%3 == 0 {
			fmt.Println("Fizz")
		} else if number%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(number)
		}

		number += 1
	}
}
