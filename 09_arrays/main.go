package main

import (
	"fmt"
)

func main()  {
	year := [3]int{20 ,30, 40}
	fmt.Println(year)
	fmt.Println(len(year))
	fmt.Println(cap(year))


	for i, v := range year{
		fmt.Printf("El index es : %v y el valor es %v\n", i, v)
	}

}
