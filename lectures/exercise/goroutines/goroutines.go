//--Summary:
//  Create a program to read a list of numbers from multiple files,
//  sum the total of each file, then sum all the totals.
//
//--Requirements:
//* Sum the numbers in each file noted in the main() function
//* Add each sum together to get a grand total for all files
//  - Print the grand total to the terminal
//* Launch a goroutine for each file
//* Report any errors to the terminal
//
//--Notes:
//* This program will need to be ran from the `lectures/exercise/goroutines`
//  directory:
//    cd lectures/exercise/goroutines
//    go run goroutines
//* The grand total for the files is 4103109
//* The data files intentionally contain invalid entries
//* stdlib packages that will come in handy:
//  - strconv: parse the numbers into integers
//  - bufio: read each line in a file
//  - os: open files
//  - io: io.EOF will indicate the end of a file
//  - time: pause the program to wait for the goroutines to finish

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"}

	var fileSum int
	var totalSum int

	numListing := func(fileName string) {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			input := scanner.Text()

			digit, err := strconv.Atoi(strings.TrimSpace(input))
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error parsing digit:", err)
			}

			fileSum += digit

			fmt.Println("Print file sum: %v", fileSum)

		}

		// Check for scanner errors
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
		}

	}

	for i := 0; i < len(files); i++ {
		go numListing(files[i])
		totalSum += fileSum
	}
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("Total sum is: %v", fileSum)
}
