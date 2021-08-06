package main

import "fmt"

type persons struct {
	name string
	last_name string
	years_old int
}

type hoobys struct {
	persons
	play string
	status bool

}

func main()  {

	persona := hoobys{persons:persons{
		name:      "Juan",
		last_name: "Moya",
		years_old: 30,
	},
		play: "basquet",
		status: true,
		}
	fmt.Println(persona)
	fmt.Println(persona.name, persona.last_name, persona.years_old, persona.play, persona.status)
}

