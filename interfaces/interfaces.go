//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Shop interface {
	available()
}

type Motorcycles string
type Cars string
type Trucks string

type GetDeveloper interface {
	getDevJob()
}

type Solidity string
type Golang string
type Rust string

func (g Golang) getDevJob() {
	fmt.Println("Golang Developer")
}

func (s Solidity) getDevJob() {
	fmt.Println("Solidity Developer")
}

func (r Rust) getDevJob() {
	fmt.Println("Rust Developer")
}

func (m Motorcycles) available() {
	// data := "Hero Honda"
	// fmt.Printf("The Motorcycle model is(%v)", data)
	fmt.Println("Motorcycle Model")
}

func (c Cars) available() {
	// data := "Baleno"
	// fmt.Printf("The car model is(%v)", data)
	fmt.Println("Car Model")
}

func (t Trucks) available() {
	// data := "Lorry"
	// fmt.Printf("The truck model is(%v)", data)
	fmt.Println("Truck Model")
}

func handleVehicle(details []Shop) {
	fmt.Println("Get vehicle details")
	for i := 0; i < len(details); i++ {
		detail := details[i]
		fmt.Printf("The vehicles detail(%v)", detail)
		detail.available()
	}
}

func iterateDeveloper(developer []GetDeveloper) {
	for i := 0; i < len(developer); i++ {
		dev := developer[i]
		fmt.Printf("The required blockchain developers for this project (%v)", dev)
		dev.getDevJob()
	}
}

func main() {

	out := []Shop{Motorcycles("Hero Honda"), Cars("Baleno"), Trucks("Lorry")}
	handleVehicle(out)

	devResult := []GetDeveloper{Solidity("Ethereum"), Golang("Go-Routines"), Rust("Cargo")}
	iterateDeveloper(devResult)

}
