//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	length int
	width  int
}

func areaCalc(rectangle Rectangle) int {
	return rectangle.length * rectangle.width
}

func perimeterCalc(rectangle Rectangle) int {
	return 2*rectangle.length + 2*rectangle.width
}

func main() {

	var rectangle = Rectangle{length: 8, width: 4}

	fmt.Println("Calculated area:", areaCalc(rectangle))
	fmt.Println("Calculated perimeter", perimeterCalc(rectangle))

}
