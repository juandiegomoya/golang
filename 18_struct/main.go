package main

import "fmt"

type person struct {
	nombre string
	apellido string
	edad int
}

func main()  {

	x := person{
		nombre: "Juan",
		apellido: "Moya",
		edad: 30,
	}
	fmt.Println(x.nombre, x.apellido, x.edad)


}
