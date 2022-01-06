package main

import "fmt"

func main()  {
	var letra string
	fmt.Println("Ingrese un numero:")
	fmt.Scanln(&letra)

	switch letra {
	case "A", "a", "E", "e", "I", "i", "O", "o", "U", "u":
		fmt.Println("es una vocal")
	default:
		fmt.Println("No es una vocal")
	}

}