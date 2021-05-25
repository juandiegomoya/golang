package main

import (
	"fmt"
)

type task struct {
	Name string
	Objetive string
	Status bool
}

func (t *task) completeTask(Status bool) {
	t.Status = Status
}

func (t *task) updateName(Name string){
	t.Name = Name
}

func (t *task) updateObjetive(Objetice string){
	t.Objetive = Objetice
}

func main()  {
	t := &task{
		Name: "Mi curso de go",
		Objetive: "Terminar el curso",
	}
	fmt.Printf("%+v\n", t)
	t.completeTask(true)
	t.updateName("otra tarea")
	t.updateObjetive("algo mas por hacer")
	fmt.Printf("%+v\n", t)
}