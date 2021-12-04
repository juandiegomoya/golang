package main

import "fmt"

type Person struct {
	name string
	age int
	
}

type Employee struct {
	Person
	id int
	status bool
}

func main()  {
	employee := Employee{}
	employee.name = "Juan Moya"
	employee.age = 30
	employee.id = 2341234
	employee.status = true
	fmt.Println(employee)

}