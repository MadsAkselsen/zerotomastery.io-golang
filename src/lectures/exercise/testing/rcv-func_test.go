package rcv

import "testing"

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

//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
func TestIncrementHealth(t *testing.T) {
	player1 := Player{
		health: 10,
		maxHealth: 10,
		energy: 5,
		maxEnergy: 5,
		name: "Mads",
	}

	player1.IncrementHealth(1)
	if player1.health > player1.maxHealth {
		t.Errorf("health should not exceed maxHealth: %v, want %v", player1.health, player1.maxHealth)
	}

	player1.IncrementHealth(-12)
	if player1.health != 0 {
		t.Errorf("health should be 0 when roming all health: %v, want %v", 0, player1.health)
	}
}



	//* If any of your  tests fail, make the necessary corrections
	//  in the copy of your rcv-func solution file.