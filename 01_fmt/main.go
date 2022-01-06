package main

import "fmt"

func main() {
	name := "juan"
	year := 30

	//Println
	fmt.Println("Un print simple")

	//Printf
	fmt.Printf("Hola %s tiene %d años\n", name, year)

	//Sprint
	message := fmt.Sprint("Hola ", name, " tiene ", year, " años\n")
	fmt.Println(message)

	//%T para saber el tipo de dato
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", year)

	//Si no se el dato uso %v
	fmt.Printf("Hola %v tiene %v años\n", name, year)

	fmt.Println("insert other name:")
	var name2 string
	fmt.Scanln(&name2)
	fmt.Println("the other name is:", name2)
}
