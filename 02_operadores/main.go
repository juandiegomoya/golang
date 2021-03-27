package main

import "fmt"

func main()  {
	x := 50
	y := 10

	//suma
	result := x + y
	fmt.Println(result)

	//resta
	result = x - y
	fmt.Println(result)

	//multiplicacion
	result = x * y
	fmt.Println(result)

	//division
	result = x / y
	fmt.Println(result)

	//modulo
	result = x % y
	fmt.Println(result)

	//incremental
	x++
	fmt.Println(x)

	//decremental
	y--
	fmt.Println(y)

	base := 15
	altura := 20
	areaTriangulo := base * altura * 2
	fmt.Println("El area de un triangulo:", areaTriangulo )
}