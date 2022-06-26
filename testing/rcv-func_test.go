//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

func newPlayer() Player {
	return Player{
		name:      "test",
		health:    100,
		maxHealth: 100,
		energy:    100,
		maxEnergy: 100,
	}
}

func TestHealth(t *testing.T) {
	player := newPlayer()
	player.statistics(100)
	if player.health > player.maxHealth {
		t.Fatalf("Health went over limit(%v)", player.health)
	}
	player.damage(player.maxHealth + 1)
	if player.health < 5 {
		t.Errorf("The Health is minimum(%v)", player.health)
	}
	if player.health > player.maxHealth {
		t.Fatalf("The health went over limit(%v)", player.health)
	}
}

func TestEnergy(t *testing.T) {
	playerEnergy := newPlayer()
	playerEnergy.statistics(100)
	if playerEnergy.energy > playerEnergy.maxEnergy {
		t.Fatalf("The energy went over limt(%v),(%v)", playerEnergy.energy, playerEnergy.maxEnergy)
	}

	if playerEnergy.energy < playerEnergy.maxEnergy {
		t.Errorf("The player is taking damage(%v)", playerEnergy.energy)
	}
}
