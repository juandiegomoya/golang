package main

import (
	"fmt"
)

func main() {
	variante()
	variante2()
	variante3()

}
func variante() {
	elemento := [5]int{1, 2, 3, 4, 5}
	porcion := elemento[1:5]
	fmt.Println(porcion)
}

func variante2() {
	elemento := make([]int, 5, 20)
	fmt.Printf("Largo %d, capacidad %d \n", len(elemento), len(elemento))

}

func variante3() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Largo %d, Capacidad %d \n", len(nums), cap(nums))
}
