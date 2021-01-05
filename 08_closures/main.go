package main

import "fmt"

func main() {
	tablaDel := 9
	miTabla := Tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Printf("9 + %d = %d \n", i, miTabla())

	}
}

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
