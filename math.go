package main

func main() {
	var (
		a = 10
		b = 20
	)

	sum := a + b
	diff := a - b
	product := a * b
	quotient := a / b
	remainder := a % b
	isEqual := a == b
	isNotEqual := a != b
	isGreater := a > b
	isLess := a < b
	isGreaterOrEqual := a >= b
	isLessOrEqual := a <= b
	println("Sum:", sum)
	println("Difference:", diff)
	println("Product:", product)
	println("Quotient:", quotient)
	println("Remainder:", remainder)
	println("Is Equal:", isEqual)
	println("Is Not Equal:", isNotEqual)
	println("Is Greater:", isGreater)
	println("Is Less:", isLess)
	println("Is Greater or Equal:", isGreaterOrEqual)
	println("Is Less or Equal:", isLessOrEqual)
	println("Arithmetic operations completed.")

	var i = 10
	i += 5
	println("Updated value of i:", i)
	i -= 3
	println("Updated value of i after subtraction:", i)

	var j = 1
	j++
	println("Value of j after increment:", j)
}
