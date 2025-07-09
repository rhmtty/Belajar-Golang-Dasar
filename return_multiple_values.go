package main

import "fmt"

func fullName() (string, string) {
	return "John", "Doe"
}

func main() {
	// firstName, lastName := fullName()
	// fmt.Println("First Name:", firstName)
	// fmt.Println("Last Name:", lastName)

	firstName, _ := fullName()
	fmt.Println("First Name (ignoring last name):", firstName)
}
