package main

import (
	"fmt"
)

func belajarArray()  {

	fmt.Println("hello dulu gak")
	// example
	var names [3]string
	names[0] = "Elvira"
	names[1] = "Elfita"
	names[2] = "Elvin"

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	var namaSiswa = [3]string {
		"ucup", "asep", "otong",
	}

	for i := 0; i < len(namaSiswa); i++ {
		fmt.Print(namaSiswa[i], " ")
	}
	// example
	var angka = [3]int {
		10,11,12,
	}

	for i := 0; i < len(angka); i++ {
		fmt.Println(angka[i])
	}

	// array banyak 
	var daftarNilai = [...]int {
		1,2,3,4,7,8,9,
	}

	for i := 0; i < len(daftarNilai); i++ {
	fmt.Print(daftarNilai[i])
	}
}