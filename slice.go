package main

import "fmt"

func main() {
	names := [...]string{"Alice", "Bob", "Charlie", "Diana", "Eve"}

	slice1 := names[1:4] // Slice from index 1 to 3
	fmt.Println(slice1)

	slice2 := names[2:] // Slice from index 2 to the end
	fmt.Println(slice2)

	slice3 := names[:3] // Slice from the start to index 2
	fmt.Println(slice3)

	slice4 := names[:] // Slice the entire array
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"  // Modifying the first element of the slice
	daySlice1[1] = "Minggu Baru" // Modifying the second element of the slice
	fmt.Println(daySlice1)
	fmt.Println(days) // The original array is also modified because slices reference the same underlying array

	daySlice2 := append(daySlice1, "Libur Baru") // Appending a new element to the slice
	daySlice2[0] = "Sabtu Lama"                  // Modifying the first element of the new slice
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days) // The original array remains unchanged for the append operation

	newSlice := make([]string, 2, 5) // Create a new slice with a capacity of 5
	newSlice[0] = "Item 1"
	newSlice[1] = "Item 2"

	fmt.Println(newSlice)
	fmt.Println("Length of newSlice:", len(newSlice))
	fmt.Println("Capacity of newSlice:", cap(newSlice))

	newSlice2 := append(newSlice, "Item 3")
	fmt.Println(newSlice2)
	fmt.Println("Length of newSlice2:", len(newSlice2))
	fmt.Println("Capacity of newSlice2:", cap(newSlice2))

	newSlice2[0] = "Item 1 Updated"
	fmt.Println(newSlice2)
	fmt.Println(newSlice) // The original slice is also modified because slices reference the same underlying array

	fromSlice := days[:]                                      // Create a slice that references the entire array
	toSlice := make([]string, len(fromSlice), cap(fromSlice)) // Create a new slice with the same length and capacity

	copy(toSlice, fromSlice)

	fmt.Println("From Slice:", fromSlice)
	fmt.Println("To Slice:", toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5, 6, 7}
	iniSlice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Initial Array:", iniArray)
	fmt.Println("Initial Slice:", iniSlice)
}
