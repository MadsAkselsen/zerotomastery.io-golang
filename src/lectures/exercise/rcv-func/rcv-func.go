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

// type Health struct {
// 	health int
// }

type Player struct {
	health int
	maxHealth int
	energy int
	maxEnergy int
	name string
}

func (player *Player) incrementHealth(healthNum int) {
	if (player.health + healthNum > player.maxHealth) {
		player.health = player.maxHealth
	} else if (player.health + healthNum < 0){
		player.health = 0
	} else {
		player.health = player.health + healthNum
	}

	fmt.Println()
	fmt.Println(">>>>", player.name, "<<<<")
	fmt.Println("HEALTH:", player.health)
	fmt.Println("MAX HEALTH:", player.maxHealth)
	fmt.Println("ENERGY:", player.energy)
	fmt.Println("MAX ENERY:", player.maxEnergy)
}

func (player *Player) incrementEnergy(energyNum int) {
	if (player.energy + energyNum > player.maxEnergy) {
		player.energy = player.maxEnergy
	} else if (player.energy + energyNum < 0){
		player.energy = 0
	} else {
		player.energy = player.energy + energyNum
	}

	fmt.Println()
	fmt.Println(">>>>", player.name, "<<<<")
	fmt.Println("HEALTH:", player.health)
	fmt.Println("MAX HEALTH:", player.maxHealth)
	fmt.Println("ENERGY:", player.energy)
	fmt.Println("MAX ENERY:", player.maxEnergy)
}

func main() {

	player1 := Player{
		health: 10,
		maxHealth: 10,
		energy: 5,
		maxEnergy: 5,
		name: "Mads",
	}

	player1.incrementHealth(-6)
	player1.incrementHealth(4)


}
