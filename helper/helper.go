package helper

import "fmt"

var version = "1.0.0"
var Application = "MyApp"

func sayGoodbye(name string) string {
	return "Goodbye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}

func Contoh() {
	sayGoodbye("John")
	fmt.Println(version)
}
