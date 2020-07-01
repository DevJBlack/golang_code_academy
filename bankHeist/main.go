package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var isHeistOn bool = true
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past th guards. Good job, but remember, this is the first step.", eludedGuards, "Needed to get above 50!")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?", "We only got", eludedGuards, "We need 50 or above, to get pass the guards")
	}

	openedVault := rand.Intn(100)

	if isHeistOn == true && openedVault >= 70 {
		fmt.Println("Grab and Go!", "We rolled a", openedVault)
	} else if isHeistOn == true && openedVault <= 69 {
		isHeistOn = false
		fmt.Println("Vault can't be opened, we are in serious trouble", "We only got a", openedVault, "We need a 70 or above to enter the vault")
	}

	leftSafely := rand.Intn(5)

	if isHeistOn == true {
		switch {
		case leftSafely == 0:
			isHeistOn = false
			fmt.Println("We didn't get pass anything...", leftSafely)
		case leftSafely == 1:
			isHeistOn = false
			fmt.Println("Welp, I didn't see that guard", "We only got a ", leftSafely)
		case leftSafely == 2:
			isHeistOn = false
			fmt.Println("Did anyone else see that Camera?", "We only got a ", leftSafely)
		case leftSafely == 3:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside...", "We only got a ", leftSafely)
		default:
			fmt.Println("We finally did it! Get to the chopper", "We got a", leftSafely)

		}
	}

	if isHeistOn == true {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("Wow we got...", amtStolen)
	}

	fmt.Println(isHeistOn)
}
