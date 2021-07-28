package main

import "fmt"

func main()  {

	x := []string{"Juan", "devops", "sre", "go"}
	y := []string{"Aleja", "Mariana", "casa", "estudiar"}

	z := [][]string{x, y}
	fmt.Println(z)
}

/*
[[Juan devops sre go] [Aleja Mariana casa estudiar]]
*/