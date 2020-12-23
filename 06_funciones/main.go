package main

import "fmt"

//la funcion lleva nombrer parametros y salida
func main() {
	fmt.Println(numero1(123))
	numero, mensaje := numero2(12)
	fmt.Println(numero)
	fmt.Println(mensaje)
}

func numero1(uno int) (h int) {
	h = uno * 100
	return
}
func numero2(num1 int) (int, string) {
	if num1 == 12 {
		return num1 * 10, "multiplicado por 10"
	} else {
		return num1 * 20, "multiplicado por 20"
	}

}

//parametros dinamicos
func dinamico() {

}
