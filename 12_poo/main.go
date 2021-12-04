package main

import (
	"fmt"
	"time"

	us "./user"
)

type pepe struct {
	us.User
}

func main() {
	u := new(pepe)
	u.AltaUsuario(1, "Mariana", time.Now(), true)
	fmt.Println(u.User)
}
