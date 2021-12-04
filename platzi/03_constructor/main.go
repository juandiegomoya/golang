package main

import "fmt"

type Task struct {
	id int
	name string
}

func NewTask(id int, name string) *Task {
	return &Task{
		id: id,
		name: name,
	}
}

func main()  {
	//1
	t := Task{}
	fmt.Println("%v\n", t)

	//2
	t2 := Task{
		id: 12341234,
		name: "Juan Moya",
	}
	fmt.Println(t2)

	//3
	t3 := new(Task)
	fmt.Println(*t3)
	t3.id = 1234
	t3.name = "Juan Moya"
	fmt.Println(*t3)

	//4
	t4 := NewTask(1231234, "Juan Moya")
	fmt.Println(*t4)
}