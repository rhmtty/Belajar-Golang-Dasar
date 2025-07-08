package main

import (
	"fmt"
)

func random() any {
	return "Hello, World!"
}

func main() {
	result := random()
	// resultString := any.(string)
	// fmt.Println(resultString)

	// resultInt := any.(int)
	// fmt.Println(resultInt) // This will panic at runtime because `any` is not an int

	switch value := result.(type) {
	case string:
		fmt.Println("String value:", value)
	case int:
		fmt.Println("Integer value:", value)
	default:
		fmt.Println("Unknown type")
	}
}
