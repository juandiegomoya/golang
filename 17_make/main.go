package main

import "fmt"

func main()  {
	x := make([]int, 10, 10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[4] = 12
	fmt.Println(x)

}
/*
[0 0 0 0 0 0 0 0 0 0]
10
10
[0 0 0 0 12 0 0 0 0 0]

*/