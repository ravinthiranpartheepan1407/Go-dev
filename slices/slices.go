//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func assemblyType(assembly string, part []string) {
	for i := 0; i < len(part); i++ {
		out := part[i]
		fmt.Println("The Part Number is:", out)
	}
}

func fruitTotal(name string, fruit []string) {
	for j := 0; j < len(fruit); j++ {
		cuts := fruit[j]
		fmt.Println("Total Cuts:", cuts)
	}
}

func main() {
	assemblyPart := []string{"First Part", "Second Part", "Third Part"}
	assemblyType("Part 1", assemblyPart)
	assemblyPart = append(assemblyPart, "Fourth Part", "Fifth Part")
	fmt.Println("Part 2", assemblyPart)
	assemblyPart3 := assemblyPart[3:]
	fmt.Println("Part 3 :", assemblyPart3)

	fruitNames := []string{"Watermelon", "Banana", "Strawberry"}
	fruitTotal("Fruit", fruitNames)
	fruitNames = append(fruitNames, "Kiwi", "Mango")
	fmt.Println("Added Fruits:", fruitNames)
	fruitCutted := fruitNames[1:]
	fmt.Println("Cutted Fruits:", fruitCutted)
}
