package main

import "fmt"

type person struct {
	name string
	last_name string
	year int
}

func main()  {

	x := person{
		name: "Juan",
		last_name: "Moya",
		year: 30,
	}
	fmt.Println(x.name, x.last_name, x.year)


}
