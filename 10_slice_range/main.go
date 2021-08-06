package main

import "fmt"

func main()  {
	simple()
}

func simple()  {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[2])
	fmt.Println(x[4])
	fmt.Println(x[6])
	fmt.Println(x[8])

	for i, v := range x {
		fmt.Println(i, " ", v)
	}

}

/*
[1 2 3 4 5 6 7 8 9]
1
3
5
7
9
0   1
1   2
2   3
3   4
4   5
5   6
6   7
7   8
8   9
*/