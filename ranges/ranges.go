package main

import "fmt"

func main() {
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
