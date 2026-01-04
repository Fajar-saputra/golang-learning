package main

import "fmt"

func main()  {

	nilaiBagus  := 90
	// nilaiLumayan  := 80
	
	// if
	if nilaiBagus > 80 {
		fmt.Println("Lulus")
	}
	
	// if else
	if nilaiBagus > 90 {
		fmt.Println("Lulusan Terbaik")
	} else {
		fmt.Println("Lulusan")
	}
	
	// if else if
	if 79 > 90 {
		fmt.Println("Lulusan Terbaik")
	} else if 79 > 80 {
		fmt.Println("Lulusan")
	} else {
		fmt.Println("Lulusan cukup")
	}

	// if expression with short statement
	numbers := [...]int {
		1,2,3,4,
	}

	if lenght := len(numbers); lenght >= 1 {
		fmt.Println("Ada isinya!")
	}

	// switch expression
	switch numbers[2] {
	case 0:
		fmt.Println("angka ")
	case 2:
		fmt.Println("angka 2")
	case 3:
		fmt.Println("angka 3")
	case 4:
		fmt.Println("angka 4")
	}


	//switch expression with short statement
	switch panjang := len(numbers); panjang > 1 {
	case true:
		fmt.Println("Let's go guys")
	case false:
		fmt.Println("I'm tired")
	}


	// tanpa exprssion
	uas := 34
	switch  {
	case uas > 90:
		fmt.Println("sangat bagus")
	case uas > 80:
		fmt.Println("bagus")
	case uas > 70:
		fmt.Println("lumayan")
	default:
		fmt.Println("hmm, ulang lagi")
	}
}