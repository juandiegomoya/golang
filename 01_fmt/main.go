package main

import "fmt"

func main()  {
	name := "juan"
	year := 30

	//Println
	fmt.Println("Un print simple")

	//Printf
	fmt.Printf("Hola %s tiene %d años\n", name, year )

	//Sprint
	message := fmt.Sprint("Hola ", name, " tiene ", year, " años\n")
	fmt.Printf(message)

	//%T
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", year)
}