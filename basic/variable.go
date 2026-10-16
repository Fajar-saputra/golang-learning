package main

// fmt untuk menampilkan sesuatu dengan method Println
import (
	"fmt"
	"reflect"
)

// Constants
	const NAMA1 = "elvira"
	const NAMA2 = "Elvira"

	var KATAGUAMAH= "B AJA"
	var (
		VIRA = "NI CIK"
		ELVIN int = 15
		FINA = true
	)

	// var bisa di deklarasikan diluar func
	var luarFunc string = "diluar func"

func  belajarVariable()  {
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
	fmt.Println(NAMA1 == NAMA2)


	// bisa pakai :=
	contohAja := "ini bisa tidak pakai var"
	fmt.Println(contohAja)


	var makan string = "nasi";
	var minum = "air putih";
	tidur := 10

	var main int
	var berlari int
	var suka bool

	// multi declare
	var a,b,c,d int = 1,2,3,4
	var part1 , part2 = 100, "Part 1"
	part3, part4 := 200, "Part 3"

	// Variable Declaration in a Block
	var (
		bagianSatu int
		bagianDua string = "bagian 2"
		bagianTiga = "bagian 3"
	)

	// name rule
	// camelCase
	var myBiniGue = "camel case"
	// PascalCase
	var MyBiniGue = "pascal case"
	// snake_case
	var my_bini_gue = "snake case"


	// tipe data
	fmt.Println(reflect.TypeOf(makan))
	fmt.Println(reflect.TypeOf(minum))
	fmt.Println(reflect.TypeOf(tidur))
	// print variabel
	fmt.Println(makan)
	fmt.Println(minum)
	fmt.Println(tidur)
	// inisiasi varibal
	main = 10
	fmt.Println(main)
	fmt.Println(berlari)
	fmt.Println(suka)
	fmt.Println(luarFunc)

	// declare
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(part1)
	fmt.Println(part2)
	fmt.Println(part3)
	fmt.Println(part4)
	fmt.Println(bagianSatu)
	fmt.Println(bagianDua)
	fmt.Println(bagianTiga)
	fmt.Println(myBiniGue)
	fmt.Println(MyBiniGue)
	fmt.Println(my_bini_gue)

	// Constants
	fmt.Println(KATAGUAMAH)
	fmt.Println(VIRA)
	fmt.Println(ELVIN)
	fmt.Println(FINA)
}