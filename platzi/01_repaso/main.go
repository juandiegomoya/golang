package main

import (
	"fmt"
	"strconv"
	"time"
)

func main()  {
	var x int
	x = 12
	y := 14
	fmt.Println(x)
	fmt.Println(y)
	myValue, err := strconv.ParseInt("14", 0 , 64)

	if err != nil {
		fmt.Println("%\n", err)
	}else {
		fmt.Println(myValue)
	}
	c := make(chan int)
	go doSomething(c)
	<-c
}

func doSomething(c chan int){
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}