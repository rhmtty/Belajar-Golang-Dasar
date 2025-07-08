package main

import "fmt"

func endApp() {
	fmt.Println("End of application")
	msg := recover()
	fmt.Println("terjadi panic:", msg)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ups Error occurred")
	}
}

func main() {
	runApp(true)
	fmt.Println("Application is running")
}
