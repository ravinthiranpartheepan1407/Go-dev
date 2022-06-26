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

type Player struct {
	name              string
	health, maxHealth uint
	energy, maxEnergy uint
}

// type Max struct {
// 	maxHealth map[Player]Name
// 	maxEnergy map[Player]Name
// }

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

	player.statistics(20)
	player.damage(5)
	player.energys(10)

}
