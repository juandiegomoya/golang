package main

import "fmt"

var num1 int
var num2 int

func main() {
	//para pedir datos de consola se realiza con scanf
	//para windows Scanln
	fmt.Println("Ingrese el numero 1")
	fmt.Scanf("%d", &num1)

	fmt.Println("Ingrese el numero 2")
	fmt.Scanf("%d", &num2)
	fmt.Printf("Los numeros ingresados son: %d y %d\n", num1, num2)
}