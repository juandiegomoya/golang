package main

import "fmt"

type vehiculo struct {
	puertas int
	color string
}

type camion struct {
	vehiculo
	cuatroRuedas bool
}

type sedan struct {
	vehiculo
	lujoso bool
}

func main()  {

	car1 := camion{vehiculo:vehiculo{
		puertas: 5,
		color: "rojo",
	},
		cuatroRuedas: true,
	}
	fmt.Println("Vehiculo #1")
	fmt.Println("\t", car1.puertas)
	fmt.Println("\t", car1.color)
	fmt.Println("\t", car1.cuatroRuedas)

	car2 := sedan{vehiculo:vehiculo{
		puertas: 4,
		color: "Negro",
	},
		lujoso: false,
	}
	fmt.Println("Vehiculo #2")
	fmt.Println("\t", car2.puertas)
	fmt.Println("\t", car2.color)
	fmt.Println("\t", car2.lujoso)
 }

