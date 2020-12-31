package main

import "fmt"

var cal func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("sumo 5 + 6 = %d \n", cal(5, 6))

	//resta

	cal = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("resta 5 - 6 = %d \n", cal(5, 6))

	cal = func(num1 int, num2 int) int {
		return num1 * num2
	}
	fmt.Printf("multiplica 5 * 6 = %d \n", cal(5, 6))

	cal = func(num1 int, num2 int) int {
		return num1 / num2
	}
	fmt.Printf("divide 5 / 6 = %d \n", cal(5, 6))

	Operaciones()
}

//Operaciones otra forma de hacer funciones anonimas
func Operaciones() {
	resultado := func() int {
		var a int = 12
		var b int = 24
		return a + b
	}
	fmt.Println(resultado())

}
