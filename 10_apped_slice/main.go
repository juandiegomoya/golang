package main

import "fmt"

func main()  {
	x := []int{1,2,3,4,5}
	fmt.Println(x)
	x = append(x , 10,11,12,13)
	fmt.Println(x)

	y := []int{20,21,22,23,24}
	x = append(x, y...)
	fmt.Println(x)
}
