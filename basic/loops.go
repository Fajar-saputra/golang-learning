package main

import "fmt"

func belajarLoops()  {
	
	count := 10

	// basic
	for i := 0; i < count; i++ {
		fmt.Print(i, " ")
	}

	// for range
	names := []string{"Elvira", "Elfita", "Elvin", "ucup", "otong", "asep", "fazil",  "fahtir", "ardom", "dimas"}

	for index, name := range names  {
		fmt.Println("index ", index, " ", name)
	}

	// break
	for i := 0; i < count; i++ {
		if i > 4 {
			break
		}

		fmt.Println("index ke ", i)
	}

	// countinue
	for i := 0; i < count; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println("index ", i)
	}

}