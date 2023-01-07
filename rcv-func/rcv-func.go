//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

// type Name string

type Fruit struct {
	types    string
	quantity int
	sold     bool
}

type Player struct {
	name              string
	health, maxHealth uint
	energy, maxEnergy uint
}

// type Max struct {
// 	maxHealth map[Player]Name
// 	maxEnergy map[Player]Name
// }

type ProgrammingSkills struct {
	name       string
	experience string
}

type Groceries struct {
	itemName string
	category string
}

// receiver function by value
func (g Groceries) shopItems(shop Groceries) Groceries {
	return Groceries{shop.itemName, g.category}
}

// receiver function by pointer
func (coded *ProgrammingSkills) programming() {
	fmt.Println(coded.name)
	if coded.name == "Golang" {
		fmt.Println("Go Programming Language")
	} else {
		fmt.Println("Learn Rust Alternatively")
	}
}

func (fruit *Fruit) createSales() any {
	if fruit.types == "Apple" {
		fmt.Println("Printed", fruit.types)
	} else {
		fmt.Println("Something else printed")
	}
	return fruit
}

func (player *Player) statistics(value uint) {
	player.health += value
	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}
	fmt.Println(player.name, "Increment", value, "Health", player.health)

}

func (player *Player) damage(value uint) {
	player.health -= value
	if player.health > player.maxHealth {
		player.health = 0
	} else {
		player.health -= value
	}
	fmt.Println(player.name, "Damage", value, "Health", player.health)
}

func (player *Player) energys(value uint) {
	player.energy += value
	if player.energy > player.maxEnergy {
		player.energy = player.maxHealth
	}

	fmt.Println(player.name, "Energy", player.energy)
}

func main() {
	player := Player{
		name:      "Ravi",
		health:    50,
		maxHealth: 100,
		energy:    250,
		maxEnergy: 500,
	}

	fruit := Fruit{
		types:    "Orange",
		quantity: 20,
		sold:     true,
	}

	fruit.createSales()

	player.statistics(20)
	player.damage(5)
	player.energys(10)

	code := ProgrammingSkills{
		name:       "Golang",
		experience: "JS or Python Previous Python Required",
	}
	code.programming()

	shopping := Groceries{"Carbonara", "Pasta"}
	fmt.Println(shopping)
}
