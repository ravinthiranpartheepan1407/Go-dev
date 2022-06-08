//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

// const (
// 	newborn 
// 	child 
// 	toddler 
// 	adult 
// ) 


func newBorn(newborn int) bool{
	return newborn == 0
}

func todDler(toddler int) int {
	if(toddler>1 && toddler<=3){
		return toddler
	}else{
		return 0
	}
}

func chIld(child int) int {
	if(child >=4 && child <=12){
		return child
	}else{
		return 0
	}
}

func adUlt(adult int) int{
	if(adult >= 18){
		return adult
	}else{
		return 0
	}
}

func main() {
	age := adUlt(22)
	fmt.Println("The age is", age)
	switch age{
	
	case age :
		// if(age >= adUlt(22)){
		// 	fmt.Println("True")
		// }else{
		// 	fmt.Println("fasle")
		// }
		fmt.Println(age)
		fmt.Println("Adult")

	default : 
		fmt.Println("Noting Detected")
	}


		
}
