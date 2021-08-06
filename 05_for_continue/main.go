package main

import "fmt"

func main()  {
	x := 1
	for{
		x++
		if x > 10{
			fmt.Println("X es mayor que 10")
			break
		}
		if x % 2 != 0 {
			fmt.Println("x es impar")
			continue
		}
		fmt.Println(x)
	}
}
