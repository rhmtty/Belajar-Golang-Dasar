package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh nol")
	}

	return nilai / pembagi, nil
}

func main() {
	hasil, err := Pembagian(10, 0)
	if err != nil {
		fmt.Println("Terjadi kesalahan:", err)
	} else {
		fmt.Println("Hasil pembagian:", hasil)
	}
}
