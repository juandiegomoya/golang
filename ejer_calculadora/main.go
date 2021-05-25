package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {}

func (calc) operate(entrada string, signo string) int {
	entradaLimpia := strings.Split(entrada, signo)
	operator1, _ := strconv.Atoi(entradaLimpia[0])
	operator2, _ := strconv.Atoi(entradaLimpia[1])

	switch signo {
	case "+":
		fmt.Println(operator1 + operator2)
		return operator1 + operator2
	case "-":
		fmt.Println(operator1 - operator2)
		return operator1 - operator2
	case "*":
		fmt.Println(operator1 * operator2)
		return operator1 * operator2
	case "/":
		fmt.Println(operator1 / operator2)
		return operator1 / operator2
	default:
		fmt.Println("no seas menso")
	}
	return 0
}

func parsear(entrada string) int {
	operator, _ := strconv.Atoi(entrada)
	return operator
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {

	entrada := leerEntrada()
	signo := leerEntrada()
	c := calc{}
	c.operate(entrada,	signo)
}