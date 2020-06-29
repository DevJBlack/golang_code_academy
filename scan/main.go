package main

import "fmt"

func main() {
	fmt.Println("How are you doing?")
	var emotion string
	fmt.Scan(&emotion)
	if emotion == "good" {
		fmt.Printf("That is great to hear that you are %v!", emotion)
	} else if emotion == "bad" {
		fmt.Printf("I am sorry to hear that you are doing %v", emotion)
	} else {
		fmt.Println("That is neither good nor bad")
	}

}
