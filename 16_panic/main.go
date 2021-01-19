package main

import (
	"log"
)

func main(){
	/*file := "test.txt"
	f, err := os.Open(file)

	defer f.Close()

	if err != nil{
		fmt.Println("Error abriendo el archivo")
		os.Exit(1)
	}*/
	ejemploPanic()
}

func ejemploPanic()  {
	defer func(){
		r := recover()
		if r != nil{
			log.Fatalf("ocurrio un error \n %v", r)
		}
	}()

	a:=1
	if a == 1{
		panic("se encontro el valor de 1")
	}

}