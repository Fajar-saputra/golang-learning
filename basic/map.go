package main

import "fmt"

func main()  {
	// example
	person := map[string]string{
		"firts name" : "Elvira",
		"last name" : "Maharani",
	}

	fmt.Println(person["firts name"])
	fmt.Println(person["last name"])

	fmt.Println(len(person))

	siswa := make(map[string]string)
	
	siswa["name"] = "Elfita"
	siswa["umur"] = "23"
	siswa["nilai"] = "89"
	siswa["tidak naik kelas"] = "Belum pernah"

	fmt.Println(siswa)
	delete(siswa, "tidak naik kelas")
	fmt.Println(siswa)
 
}