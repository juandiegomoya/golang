package main

import (
	"fmt"
)

// var tabla [5]int

// func main() {
// 	tabla[1] = 23
// 	fmt.Println(tabla)

// }

/*forma de recorrer una tabla*/
// func main() {
// 	tabla := [5]int{23, 45, 90, 75, 3}
// 	for i := 0; i < len(tabla); i++ {
// 		fmt.Println(tabla[i])
// 	}
// }
var matriz [5][7]int

func main() {
	matriz[1][6] = 1
	fmt.Print(matriz)
}
