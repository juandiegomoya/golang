package main

import "fmt"

type Employee struct {
	id int
	name string
	vacation bool
}

// forma 4 creando un a funcion

func NewEmployee (id int, name string, vacation bool) *Employee {
	return &Employee{
		id: id,
		name: name,
		vacation: vacation,
	}
}

func main()  {
	// forma 1 de construccion
	e := Employee{}
	e.id = 12
	e.name = "Juan"
	e.vacation = true
	fmt.Println(e)
	fmt.Println(e.id)
	fmt.Println(e.name)
	fmt.Println(e.vacation)

	// forma 2 de construccion

	e2 := Employee{
		id: 12,
		name: "Juand",
	}
	fmt.Println(e2)
	fmt.Println(e2.name)
	fmt.Println(e2.vacation)

	// forma 3 de construccion teniendo en en cuenta que devuelve un apuntador

	e3 := new(Employee)
	fmt.Println(e3) //cuando se usa new devuelve un apuntador
	fmt.Println(*e3)
	e3.id = 12
	e3.name = "JuanMoya"
	e3.vacation = true
	fmt.Println(e3)
	fmt.Println(*e3)

	//forma 4
	e4 := NewEmployee(12, "JuanDiego", false)
	fmt.Printf("%v", *e4)
}
