package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "John Doe",
		"address": "123 Elm Street",
	}

	fmt.Println("Name:", person["name"])
	fmt.Println("Address:", person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Go Programming"
	book["author"] = "Jane Smith"
	book["wrong"] = "Ups"

	fmt.Println(book)

	delete(book, "wrong")

	fmt.Println(book)
}
