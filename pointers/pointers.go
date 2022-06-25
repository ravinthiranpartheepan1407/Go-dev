//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type Items struct {
	items    string
	security bool
}

type Securitytags struct {
	activate   bool
	deactivate bool
}

func itemList(class Items) {
	fmt.Println(class.items)
}

func activate(act *Items) {
	itemString := make(map[string]int)
	itemString["Item A"]
	itemString["Item B"]
	itemString["Item C"]
	itemString["Item D"]
	itemList(Items{items: "Golang"})
	print(itemString)

	activation = &security
	activation := act.security
	activation = true
	act.security = activation
	result := act.security
	fmt.Println(result)
}

func deactivate(deact *Items) {
	delete(Items.items)
}

func main() {
	out := activate()
	fmt.Println(out)
}
