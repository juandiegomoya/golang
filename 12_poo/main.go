package main

import (
	"fmt"
	"time"

	us "./user"
)

type juan struct {
	us.User
}

func main() {
	u := new(juan)
	u.altaUsuario(1, "Mariana", time.Now(), true)
	fmt.Println(u.User)
}
