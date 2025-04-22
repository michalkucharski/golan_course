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

type Player struct {
	health, maxHealth uint
	energy, maxEnergy uint
	name              string
}

func (player *Player) addHealth(amount uint) {
	player.health += uint(amount)

	fmt.Println("Adding health point to player")
	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}
	fmt.Println("Actual player statistics:", player)
}

func (player *Player) addEnergy(amount uint) {
	player.energy += uint(amount)

	fmt.Println("Adding energy point to player")
	if player.energy > player.maxEnergy {
		player.energy = player.maxEnergy
	}
	fmt.Println("Actual player statistics:", player)
}

func (player *Player) applyDamage(amount uint) {

	fmt.Println("Apply damage to the player")
	if player.health-amount > player.health {
		player.health = 0
		fmt.Println("Player is dead")
	} else {
		player.health -= amount
	}

	fmt.Println("Actual player statistics:", player)
}

func (player *Player) consumingEnergy(amount uint) {

	fmt.Println("Consuming energy")

	if player.energy-amount > player.energy {
		fmt.Println("Player is exhausted")
		player.energy = 0
	} else {
		player.energy -= amount

	}

	fmt.Println("Actual player statistics:", player)
}

func main() {

	myPlayer := Player{
		name:      "Gorag",
		health:    200,
		maxHealth: 200,
		energy:    300,
		maxEnergy: 400,
	}

	myPlayer.applyDamage(34)
	myPlayer.consumingEnergy(100)
	myPlayer.addEnergy(2)
	myPlayer.addHealth(34)
	myPlayer.consumingEnergy(600)

}
