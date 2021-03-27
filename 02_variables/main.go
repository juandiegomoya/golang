package main

import (
	"fmt"
	"strconv"
)

//cuando una variable es creada por fuera puede ser usada por todo el programa
// var numero int
// var numero2 float64
// var texto string
// var status bool

func main() {
	// la variables pueden ser usadas dentro del main
	// var texto string

	// otra forma de crear una variable go identifica el tipo de dato
	// numero := 15

	// variabels encadenadas
	// var num1 , num2 num3 int

	// puedo tambien agregar texto o cualquier tipo de variable
	// num1, num2, num3 := 2, 23, "hola"

	// const pi float64 = 3.14
	// const pi2 = 3.14
	// fmt.Println(pi)
	// fmt.Println(pi2)

	// Zero values
	// var a int // 0
	// var b float64 // 0
	// var c string // ""
	// var d bool // false

	//Las variables pueden ser publicas o privadas dependiendo la primera letra en mayuscula(publica) en minuscula(privada)

	// convertir
	texto := "algo"

	var moneda int64 = 0

	texto = strconv.Itoa(int(moneda))

	fmt.Println(moneda)
	fmt.Println(texto)

}
