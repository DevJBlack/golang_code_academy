package main

import "fmt"

func fuelGauge(fuel int) {
	fmt.Println("We have this much fuel left in gallons:", fuel)
}

func calculateFuel(planet string) int {
	var fuel int

	if planet == "Venus" {
		fuel = 300000
	} else if planet == "Mercury" {
		fuel = 500000
	} else if planet == "Mars" {
		fuel = 700000
	} else {
		fuel = 0
	}
	return fuel
}

func greetPlanet(world string) {
	fmt.Printf("Welcome to your new home! %v", world)
}

func cantFly() {
	fmt.Println("We do not have the avilable fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int

	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)

	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelCost - fuelRemaining
	}
	if fuelCost > fuelRemaining {
		cantFly()
	}
	return fuelRemaining
}

func main() {
	flyToPlanet("Mars", 100)

}
