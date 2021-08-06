package main

import "fmt"

func main()  {
	foo()
	bar ("Juan")
	x := woo("Moya")
	fmt.Println(x)
	y, z := multi("Aleja" ," Ome")
	fmt.Println(y)
	fmt.Println(z)
}

//func (r receptor) identidicador(parametros) (retorno) {}
func foo()  {
	fmt.Println("Hola desde la funcion foo")
}

func bar(s string)  {
	fmt.Println("hola", s)
}

func woo(st string) string {
	return  fmt.Sprintln("Hola desde woo", st)
}

func multi(n, a string) (string, bool) {
	p := fmt.Sprint("Hola ", n, a, " desde la funcion multi")
	q := true
	return p, q
}