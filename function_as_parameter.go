package main

import "fmt"

type Filter func(string) string

func sayHelloFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "***"
	} else {
		return name
	}
}

func main() {
	sayHelloFilter("John", spamFilter) // Output: Hello John

	filter := spamFilter
	sayHelloFilter("Anjing", filter)
}
