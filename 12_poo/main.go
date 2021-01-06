package main

import (
	"fmt"
	"time"

	us "./user"
)

// type user struct {
// 	ID     int
// 	Name   string
// 	Time   time.Time
// 	Status bool
// }

// func main() {
// 	User := new(user)
// 	User.ID = 10
// 	User.Name = "Juan"
// 	fmt.Print(User)
// }
type juan struct {
	us.User
}

func main() {
	u := new(juan)
	u.altaUser(1, "juan", time.Now(), true)
	fmt.Print(u.User)
}
