package main

import "fmt"

func changeBool(tf *bool) {
	if *tf == true {
		*tf = false
	} else {
		*tf = true
	}
}

func changeFloat(cf *float64) {
	*cf += 10.27
}

func main() {
	a := false
	changeBool(&a)
	fmt.Println(a)

	var b float64
	b = 10.
	changeFloat(&b)
	fmt.Println(b)
}

//If you want to practice more with pointers and addresses:

//Create a function that changes a boolean value from a different scope. Done
//Create a function that changes a float value from a different scope.
//Find the zero value of a pointer that is initialized but hasn’t been assigned a value.
