package main

import "fmt"

func main()  {
	
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

}
