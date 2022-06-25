//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

type Calculation int

const (
	add Calculation = iota
	subtract
	multiply
	divide
)

func (calc Calculation) total(a, b int) int {
	switch calc {
	case add:
		return a + b

	case subtract:
		return a - b

	case multiply:
		return a * b

	case divide:
		return a / b
	}
	panic("Unhandled Error")
}

func main() {

	addition := Calculation(add)
	subtraction := Calculation(subtract)
	multiplication := Calculation(multiply)
	division := Calculation(divide)

	fmt.Println(addition.total(2, 2)) // = 4

	fmt.Println(subtraction.total(10, 3)) // = 7

	fmt.Println(multiplication.total(3, 3)) // = 9

	fmt.Println(division.total(100, 2)) // = 50
}
