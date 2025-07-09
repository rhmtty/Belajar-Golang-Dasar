package main

func main() {
	var names [3]string
	names[0] = "Alice"
	names[1] = "Bob"
	names[2] = "Charlie"

	println(names[0])
	println(names[1])
	println(names[2])

	values := [...]int{1, 2, 3, 4, 5}
	println(len(values))
	values[0] = 10
}
