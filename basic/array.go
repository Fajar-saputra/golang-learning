package main

import "fmt"

	func printArr(arr...string) {
		for i := 0; i < len(arr); i++ {
			fmt.Println(arr[i], " ")
		}
	}

func main()  {

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
		fmt.Println(namaSiswa)
	}
	// example
	var angka = [3]int {
		10,11,12,
	}

	for i := 0; i < len(angka); i++ {
		fmt.Println(angka[i])
	}

	// slice
	var dataWarga = [7]string {
		"kamal", "arsyad", "romi", "bintang", "egi", "yudi", "akbar",
	}



	printArr(dataWarga)
}