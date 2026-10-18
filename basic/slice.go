package main

import (
	"fmt"
	"reflect"
)

func belajarSlice()  {
	
	 dataWarga := [...]string {
			"kamal", 
			"arsyad", 
			"romi", 
			"bintang", 
			"egi", 
			"yudi", 
			"akbar",
	}

	 iniSlice := dataWarga[2:6] // salin dari index 2 sampai 6
	
	for i := 0; i < len(iniSlice); i++ {
			fmt.Print(iniSlice[i], " ")
	}

	// function slice
	fmt.Println(len(iniSlice))
	fmt.Println(cap(iniSlice))
	wargaBaru := [...]string{"fathir", "iqbal", "kelvin", "fazil"}
	fmt.Println(append(iniSlice, wargaBaru[3]))

	// make slice
	newslice := make([]string, 2,5)

	newslice[0] = "elvin"
	newslice[1] = "Elvin"

	fmt.Println(newslice)
	fmt.Println(newslice[0])
	fmt.Println(newslice[1])
	fmt.Println(len(newslice))
	fmt.Println(cap(newslice))
	
	// copy slice
	wargaPadang := dataWarga[:]
	fmt.Println(wargaPadang)
	// copySlice := make([]string,len(wargaPadang), cap(wargaPadang))
	// copy(copySlice, wargaPadang)
	// fmt.Fprintln(copySlice)


	//  array vs slice
	iniArray := [...]int {1,2,3,4,7,8,9,10}

	// iniSlice := make([]int, 4, 10)
	// iniSlice[0] = 1
	// iniSlice[1] = 2
	// iniSlice[2] = 3
	// iniSlice[3] = 4

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	fmt.Println("panjang dari wargaPadang : ", cap(wargaPadang))
	fmt.Println("type date dari wargaPadang : ", reflect.TypeOf(wargaPadang))
	fmt.Println("type date dari dataWarga : ", reflect.TypeOf(dataWarga))

	// loop slice
	dataAwal := []int{1,2,3,4,5,6,7,8,9}
	
	for _, angka := range dataAwal {
		fmt.Println("Angka : ", angka)
	}

	// cara hapus data di slice
	indexYgDiHapus := 2

	dataAwal = append(dataAwal[:indexYgDiHapus], dataAwal[indexYgDiHapus+1:]...)

	fmt.Println(dataAwal)


	// cara copy yang benar	
	mahasiswa := [...]string{"iqbal", "fazil", "riko", "adrian", "refalldo", "mulya", "gilang","yandra"}
	mahasiswaNonIT := mahasiswa[5:]

	
	fmt.Println(mahasiswa)
	fmt.Println(mahasiswaNonIT)
	fmt.Println("Type data : ", reflect.TypeOf(mahasiswa))
	fmt.Println("Type data : ", reflect.TypeOf(mahasiswaNonIT))


	sumber := []int{1,2,3,4,5,6}
	tujuan := make([]int , len(sumber))

	copy(tujuan, sumber)

	tujuan = append(tujuan[:2], tujuan[3:]...)

	fmt.Println(sumber)
	fmt.Println(tujuan)

	// copy yang benar
	items := []int{1,2,3,4,6,7,8,9}

	fmt.Println(getDataWrong(items))
}

	func getDataWrong(items []int)  []int{
		data := items
		return data[0:10]
	}

	func getDataCurrent(items []int) []int {
		// buat slice baru
		data := make([]int, 10)
		// copy isinya
		copy(data, items[10:])
		// return data
		return data
	}
