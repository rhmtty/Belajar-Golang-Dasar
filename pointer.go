package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Tangerang", "Banten", "Indonesia"}
	// address2 := address1

	// address2.City = "Jakarta"
	// fmt.Println(address1) // tidak berubah
	// fmt.Println(address2) // berubah

	address1 := Address{"Tangerang", "Banten", "Indonesia"}
	address2 := &address1

	address2.City = "Jakarta"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah
}
