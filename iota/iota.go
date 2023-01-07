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

import (
	"fmt"
)

type Calculation int

type Programming struct {
	name  string
	value int
}

const (
	solidity = iota + 2
	golang
	rust
	typescript
	python
)

const (
	add Calculation = iota
	subtract
	multiply
	divide
)

type Fruit int

const (
	apple = iota
	_
	_
	mango
	strawberry
)

func (fruit Fruit) checkModulus(data int) int {
	switch fruit {
	case apple:
		return data
	}
	panic("Unhandled Error")
}

func (p Programming) codeSkills() {
	// return []string{"Solidity", "Golang", "Rust", "Typescript", "Python"}[p]
	fmt.Println(p.name, p.value)
}

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

	fruito := Fruit(apple)
	fmt.Println(fruito.checkModulus(0))

	addition := Calculation(add)
	subtraction := Calculation(subtract)
	multiplication := Calculation(multiply)
	division := Calculation(divide)

	apples := apple
	fmt.Println(apples)

	fmt.Println(addition.total(2, 2)) // = 4

	fmt.Println(subtraction.total(10, 3)) // = 7

	fmt.Println(multiplication.total(3, 3)) // = 9

	fmt.Println(division.total(100, 2)) // = 50

	code := golang
	fmt.Println(code)

	coded := Programming{
		name:  "Golang",
		value: golang,
	}

	coded.codeSkills()
}

func checkModulus(apple int) {
	panic("unimplemented")
}
