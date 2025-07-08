package main

import "fmt"

func getCompleteName() (firstName, middleName, LastName string) {
	firstName = "John"
	middleName = "Fitzgerald"
	LastName = "Kennedy"
	return // Named return values are returned automatically
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c) // Output: John Fitzgerald Kennedy
}
