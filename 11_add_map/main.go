package main

import "fmt"

func main()  {
	m := map[string]int{
		"python": 10,
		"golang": 80,
		"javascript": 5,
	}
	fmt.Println(m)

	m["html"] = 10
	fmt.Println(m)
	/*Borrar soli si existe*/

	if v, ok := m["ruby"]; ok{
		fmt.Println("Si se encuentra", v, ok)
	}else {
		fmt.Println("No se encuentra pero lo voy a agregar", v, ok)
		m["ruby"] = 10
		fmt.Println(m)
	}
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	if v, ok := m["python"]; ok{
		fmt.Println("Si se encuentra y lo voy a borrar", v, ok)
		delete(m, "python")
		fmt.Println(m)
	}else {
		fmt.Println("No se encuentra pero lo voy a agregar", v, ok)
		m["ruby"] = 10
		fmt.Println(m)
	}

}
