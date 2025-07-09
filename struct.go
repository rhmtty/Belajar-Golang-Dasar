package main

import "fmt"

type Customer struct {
	Name, Adress string
	Age          int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var john Customer
	john.Name = "John"
	john.Adress = "Jakarta"
	john.Age = 30

	fmt.Println("Customer: ", john)

	jane := Customer{
		Name:   "Jane",
		Adress: "Bandung",
		Age:    25,
	}
	fmt.Println("Customer: ", jane)

	budi := Customer{"Budi", "Surabaya", 28}
	fmt.Println("Customer: ", budi)

	budi.sayHello("Alice")
}
