package main

func main() {
	var value32 int32 = 32768
	var value64 int64 = int64(value32)
	var value16 int16 = int16(value32)

	println("int32 value:", value32)
	println("Converted to int64:", value64)
	println("Converted to int16:", value16)

	var name string = "Alice"
	var a = name[0]
	eString := string(a)
	println(name)
	println(a)
	println("First character as string:", eString)
}
