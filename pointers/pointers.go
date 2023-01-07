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

type Fruits struct {
	types string
}

type Programming struct {
	solidity string
	golang   string
	rust     string
}

type Sold struct {
	typeSold       Fruits
	typeQuantities Quanity
	hasSold        bool
}

type Quanity struct {
	quantity int
}

func programmingType(code *Programming) string {
	return code.golang
}

func setProgramming(codes *string, new string, code *Programming) {
	*codes = new
	programmingType(code)
}

func createFruit(data *Fruits) string {
	return data.types
}

func listQuantities(qty *Quanity) int {
	return qty.quantity
}

func checkIfSold(sol *Sold) any {
	return sol
}

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
	itemString["Item A"] = 0
	itemString["Item B"] = 1
	itemString["Item C"] = 2
	itemString["Item D"] = 3
	itemList(Items{items: "Golang"})
	print(itemString)

	// activation = &security
	activation := act.security
	activation = true
	act.security = activation
	result := act.security
	fmt.Println(result)
}

// func deactivate(deact *Items) {
// 	delete(Items.items)
// }

// type FoodMenu struct {
// 	food         string
// 	price        string
// 	availability string
// }

// func checkFoodAvailability(available FoodMenu) string {
// 	return "Menu" + ":" + available.availability
// }

// func foodMenu(foodOrder FoodMenu) string {
// 	return foodOrder.food
// }

// func createTodayFoodAvailability() {
// 	create := FoodMenu{
// 		food:         "Carbonara Pasta",
// 		price:        "$15",
// 		availability: "Available",
// 	}
// 	fmt.Println(checkFoodAvailability(FoodMenu{availability: create.availability}))
// 	fmt.Println(foodMenu(FoodMenu{food: create.food}) + ":" + "  " + "Price" + "  " + create.price)
// }

func main() {
	out := activate()
	fmt.Println(out)
	programming := Programming{solidity: "Ethereum", golang: "go-routines", rust: "cargo"}
	codeData := programmingType(&programming)
	setProgramming(&codeData, "Cosmos SDK", &programming)
	fmt.Println(programming)
	fmt.Println(codeData)

	fruitData := []string{"Apple", "Orange"}
	fruitQuantity := []int{10, 25}
	fruitSold := []bool{true, false}

	outFruit := createFruit(&Fruits{types: fruitData[0]})
	outQuantity := listQuantities(&Quanity{quantity: fruitQuantity[1]})
	outSold := checkIfSold(&Sold{typeSold: Fruits{types: fruitData[0]}, typeQuantities: Quanity{quantity: fruitQuantity[1]}, hasSold: fruitSold[0]})
	fmt.Println("Fruit:", outFruit)
	fmt.Println("Quantity:", outQuantity)
	fmt.Println("Sold?", outSold)

	outFruits := createFruit(Fruits{types: fruitData[0]})
	outQuantitys := listQuantities(Quanity{quantity: fruitQuantity[1]})
	outSolds := checkIfSold(Sold{typeSold: Fruits{types: fruitData[0]}, typeQuantities: Quanity{quantity: fruitQuantity[1]}, hasSold: fruitSold[0]})
	fmt.Println("Fruit:", outFruits)
	fmt.Println("Quantity:", outQuantitys)
	fmt.Println("Sold?", outSolds)

	// createTodayFoodAvailability()
}
