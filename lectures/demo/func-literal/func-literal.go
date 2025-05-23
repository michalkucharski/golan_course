package main

import "fmt"

func add(lhs, rhs int) int {
	return lhs + rhs
}

func compute(lhs, rhs int, op func(lhs, rhs int) int) int {
	fmt.Printf("Running a computation with %v & %v\n", lhs, rhs)
	return op(lhs, rhs)
}

func main() {
	fmt.Println("2 + 3 =", compute(2, 3, add))
	fmt.Println(" 10 - 2 =", compute(10, 2, func(lhs, rhs int) int {
		return lhs - rhs
	}))

	mul := func(lhs, rhs int) int {
		fmt.Printf("multiplying %v * %v = ", lhs, rhs)
		return lhs * rhs
	}

	fmt.Println(compute(3, 4, mul))

}
