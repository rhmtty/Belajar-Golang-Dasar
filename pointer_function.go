package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndo(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{}
	ChangeCountryToIndo(&address)
	fmt.Println(address) // Output: {  Indonesia}
}
