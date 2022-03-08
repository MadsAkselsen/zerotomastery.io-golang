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

type VehicleHandler interface {
	PrintInfo()
}

type Truck string
type Car string
type Motorcycle string

type Lifts struct {
	smallLifts []Motorcycle
	standardLifts []Car
	largeLifts []Truck
}

func (t Truck) PrintInfo() {
	fmt.Println("Truck", t)
}
func (t Car) PrintInfo() {
	fmt.Println("Car", t)
}
func (t Motorcycle) PrintInfo() {
	fmt.Println("Motorcycle", t)
}

func directVehicles(vehicles []VehicleHandler) Lifts {
	fmt.Println(">>> Directing <<<")

	lifts := Lifts{
		smallLifts: make([]Motorcycle, 0,),
		standardLifts: make([]Car, 0),
		largeLifts: make([]Truck, 0),
	}

	for i := 0; i < len(vehicles); i++ {
		vehicle := vehicles[i]
		vehicle.PrintInfo()
		if truck, ok := vehicle.(Truck); ok {
			lifts.largeLifts = append(lifts.largeLifts, truck)
		} else if car, ok := vehicle.(Car); ok {
			lifts.standardLifts = append(lifts.standardLifts, car)
		} else if motorcycle, ok := vehicle.(Motorcycle); ok{
			lifts.smallLifts = append(lifts.smallLifts, Motorcycle(motorcycle))
		}
	}
	fmt.Println()
	return lifts
}



func main() {
	vehicles := []VehicleHandler{Truck("myTruck"), Car("myCar"), Motorcycle("myMotorcycle"), Truck("myTruck2")}
	lifts := directVehicles(vehicles)
	fmt.Println("Large Lifts: ",lifts.largeLifts)
	fmt.Println("Standard Lifts: ",lifts.standardLifts)
	fmt.Println("Small Lifts: ",lifts.smallLifts)
}
