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

type Time struct {
	hour, minutes, seconds int
}

type ParseError struct {
	msg   string
	input string
}

// type Developer struct {
// 	solidity string
// 	golang   string
// 	rust     string
// }

// type DeveloperTest struct {
// 	msg  string
// 	data string
// }

// func (d *DeveloperTest) getDeveloperError() string {
// 	return fmt.Sprintf("%v", d.msg, d.data)
// }

// func CheckAvailableDeveloper(data string) (Developer, error) {
// 	devData := strings.Split(data, ":")
// 	if len(devData) != 3 {
// 		return Developer{}, &DeveloperTest{"Invalid Developed Selected", data}
// 	} else {
// 		devJob, err :=
// 	}
// }

func (t *ParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

func ParseTime(input string) (Time, error) {
	components := strings.Split(input, ":")
	if len(components) != 3 {
		return Time{}, &ParseError{"Invalid Number Time Componenets", input}
	} else {
		hour, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &ParseError{fmt.Sprintf("Error Parsing Hours: %v", err), input}
		}

		minutes, err := strconv.Atoi(components[1])
		if err != nil {
			return Time{}, &ParseError{fmt.Sprintf("Error Parsing Minutes: %v", err), input}
		}
		seconds, err := strconv.Atoi(components[2])
		if err != nil {
			return Time{}, &ParseError{fmt.Sprintf("Error Parsing Seconds: %v", err), input}
		}
		if hour > 23 && hour < 0 {
			return Time{}, &ParseError{fmt.Sprintf("Hour out of range: %v", err), input}
		}
		if minutes > 59 && minutes < 0 {
			return Time{}, &ParseError{fmt.Sprintf("Minutes out of range: %v", err), input}
		}
		if seconds > 59 && seconds < 0 {
			return Time{}, &ParseError{fmt.Sprintf("Seconds out of range: %v", err), input}
		}

		return Time{hour, minutes, seconds}, nil
	}
}
