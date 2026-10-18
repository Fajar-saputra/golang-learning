package main

import "fmt"

type Person struct{
	name string
	age int
	job string
	salary int
}

type Kawan struct{
	Nama string
	Umur int
}

type Hero struct{
	Nama string
	Health int
}


func belajarStruct()  {
	var personVira, personElfita Person

	personVira.name = "Elvira"
	personVira.job = "Dosen"
	personVira.age = 20
	personVira.salary = 10000000

	personElfita = Person{name: "Elfita", job: "Pengusaha", age: 23, salary: 40000000}


	printPerson(personVira)
	printPerson(personElfita)

	// pointer
	asep := Kawan{Nama: "asep", Umur: 20}

	fmt.Println(asep)
	UlangTahun(&asep)
	fmt.Println(asep)
	
	prajurit := Hero{Nama: "Ibrahim", Health: 100}
	hasAttack(&prajurit, 34)
	fmt.Println(prajurit)
}

func printPerson(person Person){
	fmt.Println("Nama   : ", person.name)
	fmt.Println("Age    : ", person.age)
	fmt.Println("Job    : ", person.job)
	fmt.Println("Salary : ", person.salary)
	fmt.Println()
}

func UlangTahun(k *Kawan)  {
	k.Umur += 1
}

func hasAttack (h *Hero, damage int)  {
	h.Health -= damage
}