package main

import (
	"fmt"
)

type task struct {
	id int
	name string

}

type status struct {
	task
	status string
}


//serId it's a method
func (t *status) setId(id int) int{
	t.id = id
	return id
}

func (t *status) getName(name string)  string{
	t.status = name
	return name
}
func (t *status) GetStatus(status string)  {
	t.status = status
}

func main()  {
	t := status{}
	x := t.setId(1234123)
	fmt.Println(x)
	y := t.getName("juan moya")
	fmt.Println(y)
	t.GetStatus("Pending")
	fmt.Println(t.status)
}
