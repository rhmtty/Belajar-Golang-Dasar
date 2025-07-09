package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("John")
	if data == nil {
		fmt.Println("Data is nil")
	} else {
		fmt.Println("Data is not nil:", data["name"])
	}
}
