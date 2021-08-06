package main

import "fmt"

func main()  {
	x := []int {1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Printf("%T\n", x)

}

/*
[1 2 3 4 5 6 7 8 9]
9
[]int
*/