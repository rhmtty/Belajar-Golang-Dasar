package main

import "fmt"

func logging() {
	fmt.Println("Logging function executed")
}

func runApplication() {
	defer logging()
	fmt.Println("Running application")
}

func main() {
	runApplication()
}
