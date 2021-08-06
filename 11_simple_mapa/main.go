package main

import "fmt"

func main()  {
	m := map[string]int{
		"Juan": 30,
		"Aleja": 28,
		"Mariana": 1,
	}
	fmt.Println(m)
	/*Para buscar en el mapa*/
	fmt.Println(m["Aleja"])

	/*con , ok puedo validar la existencia de la busqueda*/
	v, ok := m["Mafe"]
	fmt.Println(v, ok)

	v, ok = m["Juan"]
	fmt.Println(v, ok)

	if v, ok := m["Mafe"]; ok{
		fmt.Println("La busqueda se encontro: ", v, ok)
	} else{
		fmt.Println("La busqueda se no se encuentro: ", v, ok)
	}


}
