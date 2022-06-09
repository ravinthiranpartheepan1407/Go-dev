//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must handle any number of dice, rolls, and sides
//
//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fruit(seeds int) int {
	return rand.Intn(seeds)
}

func roll(sides int) int {
	return rand.Intn(sides) + 1
}

func main() {
	rand.Seed(time.Now().UnixMicro())
	dice, sides := 2, 12
	rolls := 1

	for i := 1; i <= rolls; i++ {
		sum := 0
		for d := 1; d <= dice; d++ {
			rolled := roll(sides)
			sum += rolled
			fmt.Println("Rolled", rolled)

		}
		fmt.Println("Total Sum", sum)

		switch sum := sum; {
		case sum == 2 && dice == 2:
			fmt.Println("Snake Eyes")

		case sum == 7:
			fmt.Println("Lucky")

		case sum%2 == 0:
			fmt.Println("Even")

		case sum%2 == 1:
			fmt.Println("Odd")

		}
	}

	apple, seeds := 4, 16
	juice := 2
	for a := 0; a <= juice; a++ {
		compute := 0
		for b := 1; b <= apple; b++ {
			juiced := fruit(seeds)
			compute += juiced
			fmt.Println("Juice Quantity", juiced)
			fmt.Println("Computed Quantity", compute)

			switch compute := compute; {
			case compute == 0:
				fmt.Println("No Juice")

			case compute%2 == 0:
				fmt.Println("Juice Found")

			case compute%2 == 1:
				fmt.Println("Half Quantity found")

			}
		}
	}

}
