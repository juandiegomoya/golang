package main

import "fmt"

//la funcion lleva nombrer parametros y salida
func main() {
	fmt.Println(dinamico(12, 46))
	fmt.Println(dinamico(12, 46, 45, 45, 45))
	fmt.Println(dinamico(12))
	fmt.Println(dinamico(10, 10))
}

// func numero1(uno int) (h int) {
// 	h = uno * 100
// 	return
// }
// func numero2(num1 int) (int, string) {
// 	if num1 == 12 {
// 		return num1 * 10, "multiplicado por 10"
// 	} else {
// 		return num1 * 20, "multiplicado por 20"
// 	}

// }

//parametros dinamicos
func dinamico(num ...int) int {
	total := 0
	for Item, num := range num {
		total = total + num
		fmt.Printf("\n Item %d \n", Item)
	}
	return total
}
