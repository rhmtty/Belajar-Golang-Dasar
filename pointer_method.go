package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	john := Man{"John"}
	john.Married()
	// Output: john.Name is still "John" because the method does not modify the original
	fmt.Println(john.Name) // Output: John
}
