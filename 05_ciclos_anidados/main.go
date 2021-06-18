package main

import "fmt"

func main(){
	//for anidado
	for i := 0; i < 10; i++{
		fmt.Printf("El ciclo externo I es: %d\n", i)
		for j := 0; j <= 7; j++{
			// \t = a un espacio de tabulacion
			fmt.Printf("\t El ciclo interno J es: %d\n", j)
		}
	}
}

