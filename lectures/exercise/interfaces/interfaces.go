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

const (
	Small = iota
	Standard
	Large
)

type Lift int

type Liftpicker interface {
	Picklift() Lift
}

type Motorcycle string
type Car string
type Truck string

func (m Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(m))
}

func (t Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(t))
}

func (c Car) String() string {
	return fmt.Sprintf("Car: %v", string(c))
}

func (m Motorcycle) Picklift() Lift {
	return Small
}

func (c Car) Picklift() Lift {
	return Standard
}

func (t Truck) Picklift() Lift {
	return Large
}

func sendToLift(p Liftpicker) {
	switch p.Picklift() {
	case Small:
		fmt.Printf(" send %v to small lift\n", p)
	case Standard:
		fmt.Printf(" send %v to standard lift\n", p)
	case Large:
		fmt.Printf(" send %v to large lift\n", p)
	}
}

func main() {
	car := Car("sedan")
	truck := Truck("minery_truck")
	motorcycle := Motorcycle("ktm")

	sendToLift(car)
	sendToLift(truck)
	sendToLift(motorcycle)

}
