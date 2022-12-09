//--Summary:
//  The existing program is used to restrict access to a resource
//  based on day of the week and user role. Currently, the program
//  allows any user to access the resource. Implement the functionality
//  needed to grant resource access using any combination of `if`, `else if`,
//  and `else`.
//
//--Requirements:
//* Use the accessGranted() and accessDenied() functions to display
//  informational messages
//* Access at any time: Admin, Manager
//* Access weekends: Contractor
//* Access weekdays: Member
//* Access Mondays, Wednesdays, and Fridays: Guest

package main

import "fmt"

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func adminRole(a int) int {
	return a
}

func guestRole(g int) int {
	return g
}

func weekdays(days int) bool {
	return days <= 4
}

// func fruitTaste(taste int) int {
// 	return taste <= 0
// }

const (
	apple      = "sweet"
	mango      = "sour"
	strawberry = "sweet"
	avacado    = "butter"
)

const (
	appleQuantity      = 10
	mangoQuantity      = 20
	strawberryQuantity = 15
	avacadoQuantity    = 30
)

func check(fruitType string, quantity int) {
	if fruitType == "apple" && quantity == appleQuantity {
		fmt.Println("Apple Selected")
	} else {
		fmt.Println("Condition False")
	}
}

func fruitReturn(input string, quantity int) bool {
	if input != mango && quantity == 20 {
		fmt.Println("Printed True")
		return true
	} else {
		fmt.Println("Printed False")
		return false
	}
}

func main() {
	// The day and role. Change these to check your work.
	// if (guestRole(50) == Guest){
	// 	accessGranted()
	// 	fmt.Println("Available on Mondays, Wednesdays, Fridays")
	// } else if adminRole(60) == Admin{
	// 	fmt.Println("Available All days")
	// }else{
	// 	accessDenied()
	// }

	check("apple", appleQuantity)

	fruitReturn("mango", mangoQuantity)

	today, role := Thursday, Contractor

	if role == Admin || role == Manager {
		accessGranted()
	} else if role == Contractor && !weekdays(today) {
		accessGranted()
	} else {
		accessDenied()
	}

	// apple :=
	// fruit := apple
	// tastes := 0

	// if(fruit == "apple" && fruitTaste(tastes)){
	// 	fmt.Println("The fruit is tasty")
	// }else{
	// 	fmt.Println("The Fruit is not sweet")
	// }

}
