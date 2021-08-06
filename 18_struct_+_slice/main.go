package main

import (
	"fmt"
)

type person struct {
	name, last_name string
	list_ice []string
}
func main()  {
	p1 := person{
		name: "Juan",
		last_name: "moya",
		list_ice: []string{
			"Fresa", "Mora", "Lima", "Lulo",
		},
	}
	p2 := person{
		name: "Aleja",
		last_name: "Ome",
		list_ice: []string{
			"Coco", "Fresa", "Limon",
		},
	}
	fmt.Println(p1.name, p1.last_name)
	for i, v := range p1.list_ice{
		fmt.Println("\t", i+1, v)
	}

	fmt.Println(p2.name, p2.last_name)
	for i, v := range p2.list_ice{
		fmt.Println("\t", i+1, v)
	}
	//Otra forma de poder listar es usando un mapa
	m := map[string]person{
		p1.name: p1,
		p2.name: p2,
	}

	for _, v := range m{
		fmt.Println(v.name, v.last_name)
		for i, v := range v.list_ice{
			fmt.Println("\t", i, v)
		}
	}
}
