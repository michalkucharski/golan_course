//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these circumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func diceRoller(rolls, diceNum, sidesNum int) {

	for i := 1; i <= rolls; i++ {
		sum := 0
		roll := 0

		for j := 1; j <= diceNum; j++ {
			roll = rand.Intn(sidesNum) + 1
			sum = sum + roll

			fmt.Println("Dice #", j, "Roll #", roll)
		}

		fmt.Println("Sum of the dice roll:", sum)
		if sum == 2 && diceNum == 2 {
			fmt.Println("Snake eyes")
		} else if sum == 7 {
			fmt.Println("Lucky 7")
		} else if sum%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}

}

func main() {
	rand.Seed(time.Now().UnixMicro())

	diceRoller(3, 3, 6)
}
