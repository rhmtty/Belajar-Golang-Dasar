package main

import "fmt"

func getGoodBye(name string) string {
	return "Goodbye " + name
}

func main() {
	example := getGoodBye
	fmt.Println(example("John")) // Output: Goodbye John
}
