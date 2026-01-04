package main

import "fmt"

// example function basic
func helloWorld()  {
	fmt.Println("Hello World!")
}

// function with parameter
func kataKataHariIni(kataKata string)  {
	fmt.Println("Kata kata hari ini masssen: ", kataKata)
}

// functoin with return
func add(x int, y int) int {
	return  x + y;
}

// function with parameter and return mutliple value
func getNameKawan() (teman1, teman2, teman3 string) {
	teman1 = "Elvira"
	teman2 = "Elfita"
	teman3 = "Elvin"
	
	return teman1, teman2, teman3
}

// variadic function
func sumAll(numbers...int) int {
	total := 0;

	for _, number := range numbers{
		total += number
	}

	return total
}

func main()  {
	helloWorld()
	kataKataHariIni("tinggalkan semua harapan, kecuali golang!!")
	add(10, 20)

	kawanPertama, kawanKedua, kawanKetiga := getNameKawan() 
	fmt.Println(kawanPertama," ", kawanKedua," ", kawanKetiga)

	total := sumAll(1,2,3,4,4432,324,34,213,42,10)
	fmt.Println(total)
	nilai := []int{80,70,88,87,82,91}
	totalNilai := sumAll(nilai...)
	fmt.Println(totalNilai)
}