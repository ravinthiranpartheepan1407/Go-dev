//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Shopping struct {
	name  string
	price int
}

func shoppingList(shop [4]Shopping) {
	for i := 0; i < len(shop); i++ {
		shops := shop[i]
		fmt.Println(shops.name)
	}
}

func setPrice(pricez [4]Shopping) {
	for j := 0; j < len(pricez); j++ {
		prices := pricez[j]
		fmt.Println(prices.price)
	}
}

type Fruits struct {
	types   string
	quanity int
}

func validateFruit(data [3]Fruits) {
	fmt.Println(data)
}

func sum(total [4]Shopping) {
	fmt.Println(total)
}

func total(class Shopping) {
	fmt.Println(class.price + class.price)
}

type Cart struct {
	name      string
	available bool
}

type Programming struct {
	solidity string
	golang   string
	rust     string
}

func programmingCheck(programmingcheck [7]Programming) {
	for i := 0; i < len(programmingcheck); i++ {
		programmingchecks := programmingcheck[i]
		fmt.Println(programmingchecks)
	}
}

func cartAvailable(cart [4]Cart) {
	for k := 0; k < len(cart); k++ {
		carts := cart[k]
		fmt.Println(carts)
	}
}

type IcecreamMachine struct {
	iceCreamName string
	iceCreamType string
}

func getChocolateIceCream(icecream [4]IcecreamMachine) {
	for i := 0; i < len(icecream); i++ {
		icecreams := icecream[i]
		getChocolate := icecreams.iceCreamName
		if getChocolate == "Chocolate" {
			fmt.Println("Got Chocolcate Ice Cream")
		} else {
			fmt.Println("Not able to find the chocolate flavor")
		}
	}
	fmt.Println(icecream)
}

func main() {
	shops := [...]Shopping{
		{name: "producta"},
		{name: "productb"},
		{name: "productc"},
		{name: "productd"},
	}

	programming := [...]Programming{
		{solidity: "Ethereum"},
		{golang: "Go-Routines"},
		{rust: "Cargo"},
		{solidity: "Polygon"},
		{golang: "Concurrent"},
		{rust: "Anchor"},
		{solidity: "Solana"},
	}

	fruit := [...]Fruits{
		{types: "Mango", quanity: 20},
		{types: "Apple", quanity: 10},
		{types: "Strawberry", quanity: 30},
	}

	validateFruit(fruit)

	prices := [...]Shopping{
		{price: 25},
		{price: 35},
		{price: 45},
		{price: 55},
	}

	carts := [...]Cart{
		{name: "Master", available: true},
		{name: "Visa", available: false},
		{name: "Mastero", available: true},
		{name: "default", available: true},
	}

	shoppingList(shops)
	setPrice(prices)
	cartAvailable(carts)
	fmt.Println(carts[1])

	programmingCheck(programming)
	fmt.Println(programming[2])

	total(Shopping{price: 25})

	icecreamMetadata := [...]IcecreamMachine{
		{iceCreamName: "Vanilla", iceCreamType: "vanilla Flavor"},
		{iceCreamName: "Butterscotch", iceCreamType: "Butterscotch Flavor"},
		{iceCreamName: "Chocolate", iceCreamType: "Chocolate Flavor"},
		{iceCreamName: "Mango", iceCreamType: "Mango Flavor"},
	}

	getChocolateIceCream(icecreamMetadata)

}
