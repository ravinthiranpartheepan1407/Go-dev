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

	integer := 1

	for i := 0; i < 50; i++ {
		fizz := integer + i
		if fizz%3 == 0 {
			fmt.Println("Fizz")
		} else if fizz%5 == 0 {
			fmt.Println("Buzz")
		} else if fizz%3 == 0 && fizz%5 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println("There is an error")
		}
	}

	fruit := 1

	for j := 0; j < 50; j++ {
		print := fruit + j
		if print%2 == 0 {
			fmt.Println("Even Fruits")
		} else if print%4 == 0 {
			fmt.Println("Even 4 Fruits")
		} else {
			fmt.Println("Fruit out of bound")
		}
	}

	count := 0
	seconds := 0
	for i := 0; i < 150; i++ {
		message := count + 1
		second := seconds + 1
		if message == second {
			fmt.Println("Yeah I have apologized only for", i)
		} else {
			fmt.Println("Okay I will get a slap from you!")
		}
	}
}
