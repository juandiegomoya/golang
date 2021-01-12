package main

import "fmt"

type humano interface {
	respirar()
	comer()
	pensar()
	sexo() string
}

type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
}

type vegetal interface {
	ClasificacionVegetal()
}

type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
}

type mujer struct {
	hombre
}

func (h *hombre) respirar() { h.respirando = true }
func (h *hombre) comer()    { h.comiendo = true }
func (h *hombre) pensar()   { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHombre == true {
		return "Hombre"
	} else {
		return "Mujer"
	}
}

//HumanoRespirando es la funcion que indica que esta respirando
func HumanoRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un@ %s, y ya estoy respirando \n", hu.sexo())

}

type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
}

func (p *perro) respirar()         { p.respirando = true }
func (p *perro) comer()            { p.comiendo = true }
func (p *perro) EsCarnivoro() bool { return p.carnivoro }

//AnimalesRespirar funcion de respirar
func AnimalesRespirar(an animal) {
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

//AnimalesCarnivoros de tipo de alimentacion
func AnimalesCarnivoros(an animal) int {
	if an.EsCarnivoro() == true {
		return 1
	}
	return 0
}

func main() {
	Pedro := new(hombre)
	Pedro.esHombre = true
	HumanoRespirando(Pedro)

	Maria := new(mujer)
	HumanoRespirando(Maria)

	totalCarnivoros := 0
	Yago := new(perro)
	Yago.carnivoro = true
	AnimalesRespirar(Yago)
	totalCarnivoros = +AnimalesCarnivoros(Yago)
	fmt.Printf("Total carnivoros %d \n", totalCarnivoros)
}
