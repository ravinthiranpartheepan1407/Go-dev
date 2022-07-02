//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	hello = "Hello"
	hi    = "Hi"
)

type Book struct {
	hellos string
	his    string
}

// func (t *Book) Books() string {
// 	return fmt.Sprintf("%v %v", t.hellos, t.his)
// }

func readBooks(class Book) {
	fmt.Println(class.hellos, class.his)
}

func main() {

	r := bufio.NewReader(os.Stdin)

	sum := 0

	readBooks(Book{hellos: "Hello", his: "hi"})

	scanner := bufio.NewScanner(os.Stdin)
	lines := 0
	commands := 0

	scanner.Scan()
	for scanner.Scan() {
		if strings.ToUpper(scanner.Text()) == "Q" {
			break
		} else {
			text := strings.TrimSpace(scanner.Text())

			switch text {
			case hello:
				commands += 1
				fmt.Println("Hello")
			case hi:
				commands += 1
				fmt.Println("Hi")
			}

			if text != "" {
				lines += 1
			}

		}
	}

	fmt.Println("Lines \n", lines)
	fmt.Println("Commands \n", commands)

	for {
		input, inputErr := r.ReadString(' ')
		n := strings.TrimSpace(input)
		if n == "" {
			continue
		}
		num, conErr := strconv.Atoi(n)
		if conErr != nil {
			fmt.Println(conErr)
		} else {
			sum += num
		}

		if inputErr == io.EOF {
			break
		}
		if inputErr != nil {
			fmt.Println("Error Reading:", inputErr)
		}
	}
	fmt.Printf("Sum: %v", sum)

	readBooks := bufio.NewReader(os.Stdin)

	character := os.Stdin

	for {
		inputBook, inputBookError := readBooks.ReadString(' ')
		m := strings.TrimSpace(inputBook)
		if m == "" {
			continue
		}

		text, convertError := fmt.Println(character)

		if convertError != nil {
			fmt.Println(convertError)
		} else {
			fmt.Println(text)
		}
		if inputBookError == io.EOF {
			break
		}

		if inputBookError != nil {
			fmt.Println("Error Reading Books: %v", inputBookError)
		}

	}

	fmt.Println("The Book are: %v", character)

}
