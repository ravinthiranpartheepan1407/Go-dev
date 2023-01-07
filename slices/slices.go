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

// Make used to reallocate size

package main

import "fmt"

type Part string

type Fruit struct {
	quanity int
	healthy string
}

type Player struct {
	health    int
	maxHealth int
}

type BlockchainDevPath struct {
	path1 string
	path2 string
	path3 string
}

// segmentSlices := make([] string, 4)

func blockchainPath(route string, segment []string) {
	for i := 0; i < len(segment); i++ {
		pathway := segment[i]
		fmt.Println(pathway)
	}
}

func printHealth(healthBar int, input []int) int {
	for i := 0; i < len(input); i++ {
		output := input[i]
		fmt.Println("Printed output", output)
	}
	return healthBar
}

func print(value string, data []int) {
	for i := 0; i < len(data); i++ {
		res := data[i]
		fmt.Println("Printed data", res)
	}
}

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

func checkShoes(shoe []string) {
	for i := 0; i < len(shoe); i++ {
		shoes := shoe[i]
		// if shoe[0] == "GroupA" && shoe[2] == "GroupC" {
		// 	fmt.Println("Shoes for players with foot size 6 and 7")
		// } else if shoe[1] == "GroupB" {
		// 	fmt.Println("Shoes for players with foot size: 9")
		// } else {
		// 	fmt.Println(shoes)
		// }
		fmt.Println(shoes)
	}
}

func main() {
	assemblyPart := []string{"First Part", "Second Part", "Third Part"}
	assemblyType("Part 1", assemblyPart)
	assemblyPart = append(assemblyPart, "Fourth Part", "Fifth Part")
	fmt.Println("Part 2", assemblyPart)
	assemblyPart3 := assemblyPart[3:]
	fmt.Println("Part 3 :", assemblyPart3)

	printHealthBar := []int{95, 100, 95, 58}
	printHealth(10, printHealthBar[:2])
	printSecondHealth := printHealthBar[1:3]
	printHealth(20, printSecondHealth)

	fruitPrint := []int{10, 20, 30, 40, 50}
	print("Mango", fruitPrint)

	fruitNames := []string{"Watermelon", "Banana", "Strawberry"}
	fruitTotal("Fruit", fruitNames)
	fruitNames = append(fruitNames, "Kiwi", "Mango")
	fmt.Println("Added Fruits:", fruitNames)
	fruitCutted := fruitNames[1:]
	fmt.Println("Cutted Fruits:", fruitCutted)

	blockPathway := []string{"Solidity", "Golang", "Rust"}
	blockchainPath("Path A:", blockPathway[:1])

	blockPathwayAppend := append(blockPathway, "Typescript", "Javascript", "Python")
	blockchainPath("Overall Pathway:", blockPathwayAppend)

	playersFootSize := []string{"GroupA", "GroupB", "GroupC"}

	checkShoes(playersFootSize[:1])

}
