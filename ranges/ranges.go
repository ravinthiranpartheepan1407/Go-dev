package main

import "fmt"

func rangeCheck() {
	data := []string{"Apple", "Mango", "Strawberry"}
	for i, values := range data {
		fmt.Println(i, values)
		for _, split := range values {
			fmt.Printf("%q", split)
		}
	}
}

func programmingCheck() {
	programming := []string{"Solidity", "Golang", "Rust"}
	for i, split := range programming {
		fmt.Println(i, split)
		for _, check := range split {
			fmt.Printf("%q \n", check)
		}
	}
}

func main() {

	rangeCheck()

	programmingCheck()

	// %q split double quotes into runes
	fruit := []string{"Watermelon", "Banana", "Kiwi"}
	fmt.Println(fruit)
	for index, element := range fruit {
		fmt.Println(index, element)
		for _, count := range element {
			fmt.Printf(" %q \n", count)
		}
	}

	movies := []string{"Terminator", "Horror", "DC"}
	for movieIndex, content := range movies {
		fmt.Println(movieIndex, content)
		for _, splits := range content {
			fmt.Printf("%q", splits)
		}
	}

	movies = append(movies, "Marvel")
	fmt.Println(movies)
	moviesWon := movies[0:1]
	fmt.Println("Blockbusters:", moviesWon)

}
