//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

type LineRune func(line string)

func LineIteration(lines []string, call LineRune) {
	for i := 0; i < len(lines); i++ {
		call(lines[i])
	}
}

type ProgrammingLine func(code string)

func ProgrammingIterate(codes []string, codez ProgrammingLine) {
	for c := 0; c < len(codes); c++ {
		codez(codes[c])
	}
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	codeline := 0
	compilation := 0

	coded := []string{
		"Solidity Fifty codeline,",
		"Golang zero compilation,",
		"Rust 50,",
	}

	letters := 0
	digits := 0
	spaces := 0
	punctuations := 0

	programmingFunction := func(code string) {
		for _, codes := range code {
			if unicode.IsLetter(codes) {
				codeline += 1
			}
			if unicode.IsNumber(codes) {
				compilation += 1
			}
		}
	}

	ProgrammingIterate(coded, programmingFunction)
	fmt.Println("Letters", codeline)
	fmt.Println("Numbers", compilation)

	lineFunction := func(line string) {
		for _, liners := range line {
			if unicode.IsLetter(liners) {
				letters += 1
			}
			if unicode.IsNumber(liners) {
				digits += 1
			}
			if unicode.IsSpace(liners) {
				spaces += 1
			}
			if unicode.IsPunct(liners) {
				punctuations += 1
			}
		}
	}
	LineIteration(lines, lineFunction)
	fmt.Println("Letters", letters)
	fmt.Println("Digits", digits)
	fmt.Println("Spaces", spaces)
	fmt.Println("Punctuations", punctuations)

}
