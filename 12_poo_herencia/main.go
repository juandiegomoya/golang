package main

import "fmt"

type Person struct {
	name string
	age int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	status bool
}

func main()  {
	newPerson := FullTimeEmployee{}
	fmt.Println(newPerson)
	newPerson.id = 12
	newPerson.name = "Otto"
	newPerson.Person.age = 40
	newPerson.status = true
	fmt.Println(newPerson)

}
