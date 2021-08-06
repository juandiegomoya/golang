package main

import "fmt"

func main()  {
	p := struct {
		name, last_name string
		years_old int
	}{
		name: "Juan",
		last_name: "Moya",
		years_old: 30,
	}
	fmt.Println(p)
	fmt.Println(p.name)
	fmt.Println(p.last_name)
	fmt.Println(p.years_old)
}
