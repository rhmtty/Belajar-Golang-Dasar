package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Counter: ", counter)
	// 	counter++
	// }

	for i := 1; i <= 10; i++ {
		fmt.Println("Counter: ", i)
	}

	names := []string{"John", "Jane", "Doe"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Printf("Index: %d, Name: %s\n", index, name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
