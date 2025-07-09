package main

import "fmt"

func main() {
	name := "John Doe"

	if name == "John" {
		fmt.Println("Hello, John!")
	} else if name == "Jane" {
		fmt.Println("Hello, Jane!")
	} else {
		fmt.Println("Hello, stranger!")
	}

	if length := len(name); length > 5 {
		fmt.Println("Your name is quite long!")
	} else {
		fmt.Println("Your name is short.")
	}
}
