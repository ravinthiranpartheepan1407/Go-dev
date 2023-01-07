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

const (
	experiencedWarriors = 50000
)

func personName() string {
	return "ravi"
}

func ignore(f, g, _ int) int {
	return f * g
}

func add(a, b int) int {
	return a + b
}

func mul(c, d int) int {
	return c * d
}

func srq(e int) int {
	return e * e
}

func message() {
	fmt.Println("Welcome")
}

func dataOut() string {
	return "Hello"
}

func fruit(data string) string {
	return data
}

func absolute(aa, bb int) int {
	return aa + bb
}

func block(web3, web4 int) int {
	return web3 + web4
}

func totalWarriors(youngWarriors int) int {
	return youngWarriors + experiencedWarriors
}

func youngWarriorsClan(youngWarriors int) int {
	return youngWarriors
}
func main() {

	youngSullys := youngWarriorsClan(35000)
	totalNaviWarriros := totalWarriors()
	fmt.Println("Total warriors sully has so far:", totalNaviWarriros)
	fmt.Println("Young warriors in sully's clan:", youngSullys)

	data := dataOut()
	fmt.Println("Welcome", personName())

	ignoreOut := ignore(5, 6, 7)
	fmt.Println("Ignored", ignoreOut)

	addNum := add(5, 6)
	sqrNum := srq(6)
	mulNum := mul(7, 7)

	addOutput := addNum
	sqrmulOut := sqrNum + mulNum
	fmt.Println("The output is", addOutput)
	fmt.Println("The Squared Multiplied Number is", sqrmulOut)
	fmt.Println("The message is ", data)

	printAbsolute := absolute(5, 6)
	fmt.Println("The printed absolute value", printAbsolute)

	print := fruit("Apple")
	fmt.Println("The printed fruit name", print)

	fmt.Println("The printed number", block(4, 3))

}
