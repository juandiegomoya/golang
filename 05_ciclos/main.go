package main

import "fmt"

//Ciclo for de ejemplo.
// func main() {
// 	i := 1
// 	for i < 10 {
// 		fmt.Println(i)
// 		i++
// 	}
// }

// Segundo ejemplo de como usar un for
// func main() {
// 	for i := 0; i <= 10; i++ {
// 		fmt.Println(i)
// 	}
// }

// tercer ejemplo
// func main() {
// 	num := 0
// 	for {
// 		fmt.Println("continuo...")
// 		fmt.Println("Encuentra el numero secreto!!!")
// 		fmt.Scanln(&num)
// 		if num == 77 {
// 			fmt.Println("encontraste el numero secreto!!!")
// 			break
// 		}
// 	}
// }

// el Printf sedigo %d pone la variable en el texto para asi formarlo
// func main() {
// 	var num = 0
// 	for num < 10 {
// 		fmt.Printf("el valor de i es: %d", num)
// 		if num == 6 {
// 			fmt.Println("	multiplico por 2 ")
// 			num = num * 2
// 			continue
// 		}
// 		fmt.Println("	aqui pase...")
// 		num++
// 	}

// }

func main() {
	var i int = 0
RUTINAS:
	for i < 10 {
		if i == 7 {
			i = i + 1
			fmt.Println("Voy a rutina sumando 2 a i")
			goto RUTINAS
		}
		fmt.Printf("el valor de i es: %d \n", i)
		i++
	}
}
