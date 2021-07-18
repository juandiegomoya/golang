package main

import "fmt"

func main(){
	i := 0
	for {
		if i > 7 {
			fmt.Println("Aqui va el break")
			break
		}
		fmt.Println(i)
		i++
	}
}
