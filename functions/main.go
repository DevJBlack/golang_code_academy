package main

import "fmt"

func queryDatabase(query string) string {
	// Creating a varaiable, can also do result := ""
	var result string
	// Calling the function connectDatabase() to print "Connecting to the database"
	connectDatabase()
	// We are defering this function so it runs at the very end of the function
	defer disconnectDatabase()
	// We are passing in this string from our main function
	if query == "SELECT * FROM coolTable;" {
		result = "NAME|DOB\nVincent Van Gogh|March 30, 1853"
	}
	fmt.Println(result)
	return result
}

func connectDatabase() {
	fmt.Println("Connecting to the database.")
}

func disconnectDatabase() {
	fmt.Println("Disconnecting from the database.")
}

func main() {
	queryDatabase("SELECT * FROM coolTable;")
}
