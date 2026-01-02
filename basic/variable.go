package main

// fmt untuk menampilkan sesuatu dengan method Println
import "fmt"

func  main()  {
	fmt.Println("Hello, Golang 🚀")

	// tipe data int
	fmt.Println("Angka random", 1)

	// tipe data string
	fmt.Println("yah ini string")
	fmt.Println("yah ini string"[1])
	fmt.Println(len("yah ini string"))

	// variabel
	var nama = "Fajar saputra"
	fmt.Println(nama)

	nama = "nama diganti"
	fmt.Println(nama)

	jurusan := "Teknik Informatika"
	fmt.Println(jurusan)

	var (
		namaSiswa = "Elvira Maharani"
		kelasSiswa = 12
		umurSiswa = 18
	)

	fmt.Println(namaSiswa)
	fmt.Println(kelasSiswa)
	fmt.Println(umurSiswa)

	const statusNikah = false
	fmt.Println(statusNikah)

	const (
		senjata = "shogun"
		level = 38
		damage = 2899
	)

	fmt.Println(senjata)
	fmt.Println(level)
	fmt.Println(damage)

	// type declarations
	type NoKtp string
	var ktpElvira NoKtp= "123123123"
	fmt.Println(ktpElvira) 
	fmt.Println(NoKtp("3211321321")) 
	fmt.Println()

	// perbandingan
	const nama1 = "elvira"
	const nama2 = "Elvira"
	fmt.Println(nama1 == nama2)


	// bisa pakai :=
	contohAja := "ini bisa tidak pakai var"
	fmt.Println(contohAja)
}