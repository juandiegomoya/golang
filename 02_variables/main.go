package main

import (
	"fmt"
	"strconv"
)

//cuando una variable es creada por fuera puede ser usada por todo el programa
// var numero int
// var texto string
// var status bool

func main() {
	// la variables solo es usada dentro del main
	// var texto string

	// otra forma de crear una variable go identifica el tipo de dato
	// numero := 15

	// variabels encadenadas
	// var num1 , num2 num3 int

	// puedo tambien agregar texto o cualquier tipo de variable
	// num1, num2, num3 := 2, 23, "hola"

	//NOTA: go inicializa las variables en 0 o cadenas vacias o si es un bool en false
	//Las variables pueden ser publicas o privadas dependiendo la primera letra en mayuscula(publica) en minuscula(privada)

	// convertir
	texto := "algo"

	var moneda int64 = 0

	texto = strconv.Itoa(int(moneda))

	fmt.Println(moneda)
	fmt.Println(texto)

}
