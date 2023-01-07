//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
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

type ProgrammingLanguages struct {
	solidity string
	golang   string
	rust     string
}

type Rectangle struct {
	length int
	width  int
}

type Fruit struct {
	types    string
	quantity int
}

func programmingPrint(code ProgrammingLanguages) {
	fmt.Println(code.solidity, code.golang, code.rust)
}

func print(data Fruit) {
	fmt.Println(data.types, data.quantity)
}

func rectangle(class Rectangle) {
	fmt.Println(class.length * class.width)

}

func perimeter(class Rectangle) {
	fmt.Println(class.length + class.width)
}

type Name struct {
	first string
	last  string
}

func insertName(class Name) {
	fmt.Println(class.first, class.last)
}

func main() {
	var Calculates Rectangle
	Calculates.length = 5
	Calculates.width = 5
	calc := Calculates.length * Calculates.width
	fmt.Println(calc)

	var Sum Rectangle
	Sum.length = 5
	Sum.width = 7
	calcPerimeter := Sum.length + Sum.width
	fmt.Println(calcPerimeter)

	rectangle(Rectangle{length: 5, width: 5})
	perimeter((Rectangle{length: 7, width: 7}))

	insertName(Name{first: "Suren", last: "Ravi"})

	print(Fruit{types: "Mango", quantity: 20})

	programmingPrint(ProgrammingLanguages{solidity: "ETH", golang: "Go-Routines", rust: "Cargo"})
}
