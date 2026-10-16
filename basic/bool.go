package main

import "fmt"

func belajarBool()  {

	const (
		nilaiAkhir int = 90
		absen int = 3
   
		tuntasNilai bool = nilaiAkhir > 80
		absenAman bool= absen < 4
   
		lulus = tuntasNilai && absenAman
	)

	fmt.Println(lulus)
}