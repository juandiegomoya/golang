package main

import "fmt"

// var estado bool

// func main() {
// 	// Podemos incluir una variable en la misma linea de la sentencia
// 	if estado = false; estado == true {
// 		fmt.Println("Es verdadero")
// 	// con else if podemos hacer anidacion de condiciones.
// 	} else {
// 		fmt.Println("Es falso")
// 	}
// }

func main() {
	switch num := 10; num {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("el numero es mayor a 5")
	}
}
