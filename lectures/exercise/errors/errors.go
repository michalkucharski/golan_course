//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type TimeConvert struct {
	hour, minute, second int
}

type ParseError struct {
	msg   string
	input string
}

func (er *ParseError) Error() string {
	return fmt.Sprintf("%v:  %v:", er.msg, er.input)
}

func convStrToTime(strTime string) (TimeConvert, error) {
	timeList := strings.Split(strTime, ":")
	fmt.Println("Spliting time", timeList)

	if len(timeList) != 3 {
		return TimeConvert{}, &ParseError{"Invalid number of time components", strTime}

	} else {
		hour, err := strconv.Atoi(timeList[0])
		if err != nil {
			return TimeConvert{}, &ParseError{fmt.Sprintf("Error parsing hour %v", err), strTime}
		}

		minute, err := strconv.Atoi(timeList[1])
		if err != nil {
			return TimeConvert{0, 0, 0}, &ParseError{fmt.Sprintf("Error parsing minute %v", err), strTime}
		}

		second, err := strconv.Atoi(timeList[2])
		if err != nil {
			return TimeConvert{0, 0, 0}, &ParseError{fmt.Sprintf("Error parsing second %v", err), strTime}
		}

		if hour < 0 || hour > 23 {
			return TimeConvert{0, 0, 0}, &ParseError{fmt.Sprintf("Error incorrect value of hour %v", err), strTime}
		}

		if minute < 0 || minute > 59 {
			return TimeConvert{0, 0, 0}, &ParseError{fmt.Sprintf("Error incorrect value of minute %v", err), strTime}
		}

		if minute < 0 || minute > 59 {
			return TimeConvert{0, 0, 0}, &ParseError{fmt.Sprintf("Error incorrect value of second %v", err), strTime}
		}

		return TimeConvert{hour, minute, second}, nil

	}

}
