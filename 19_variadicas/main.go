package main

import "fmt"

func sumNum(values ...int) int {
	total := 0
	for _, x := range values {
		total += x
	}
	return total
}

func listName(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func main() {
	fmt.Println(sumNum(1, 2, 3, 4, 5, 6, 7, 8, 9))
	listName("Juan", "Aleja", "Mafe", "Mariana")
}
