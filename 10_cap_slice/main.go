package main

import "fmt"

func main()  {
	x := []int {1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++{
		fmt.Println(i, x[i])
	}
}
/*
9
9
0 1
1 2
2 3
3 4
4 5
5 6
6 7
7 8
8 9
*/