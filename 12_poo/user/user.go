package user

import "time"

//User is a new struct
type User struct {
	ID     int
	Name   string
	Date   time.Time
	Status bool
}

func (aqui *User) AltaUsuario(id int, name string, date time.Time, status bool) {
	aqui.ID = id
	aqui.Name = name
	aqui.Date = date
	aqui.Status = status
}
