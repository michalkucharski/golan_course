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

package testing

import "testing"

func TestApplyDamage(t *testing.T) {

	myPlayer := Player{
		name:      "Gorag",
		health:    200,
		maxHealth: 200,
		energy:    300,
		maxEnergy: 400,
	}

	myPlayer.ApplyDamage(100)
	if myPlayer.health != 100 {
		t.Errorf("incorrect health level: is %v should be %v", myPlayer.health, 100)
	}

	myPlayer.ApplyDamage(400)
	if myPlayer.health != 0 {
		t.Errorf("incorrect health level: is %v should be %v", myPlayer.health, 0)
	}

}

func TestAddHealth(t *testing.T) {

	myPlayer := Player{
		name:      "Gorag",
		health:    200,
		maxHealth: 200,
		energy:    300,
		maxEnergy: 400,
	}

	myPlayer.ApplyDamage(100)

	myPlayer.AddHealth(50)

	if myPlayer.health != 150 {
		t.Errorf("Incorrect health level: is %v, should be %v", myPlayer.health, 150)
	}

	myPlayer.AddHealth(300)

	if myPlayer.health > myPlayer.maxHealth {
		t.Errorf("The health level cant be greater than max level and is %v", myPlayer.health)
	}
}

func TestAddEnergy(t *testing.T) {

	myPlayer := Player{
		name:      "Gorag",
		health:    200,
		maxHealth: 200,
		energy:    300,
		maxEnergy: 400,
	}

	myPlayer.ConsumingEnergy(200)

	myPlayer.AddEnergy(70)

	if myPlayer.energy != 170 {
		t.Errorf("Incorrect energy level: is %v, should be %v", myPlayer.energy, 170)
	}

	myPlayer.AddEnergy(300)

	if myPlayer.energy > myPlayer.maxEnergy {
		t.Errorf("The energy level cant be greater than max level and is %v", myPlayer.energy)
	}
}

func TestConsumingEnergy(t *testing.T) {

	myPlayer := Player{
		name:      "Gorag",
		health:    200,
		maxHealth: 200,
		energy:    300,
		maxEnergy: 400,
	}

	myPlayer.ConsumingEnergy(100)
	if myPlayer.energy != 200 {
		t.Errorf("incorrect energy level: is %v should be %v", myPlayer.energy, 200)
	}

	myPlayer.ConsumingEnergy(400)
	if myPlayer.energy != 0 {
		t.Errorf("incorrect healenergyth level: is %v should be %v", myPlayer.energy, 0)
	}

}
