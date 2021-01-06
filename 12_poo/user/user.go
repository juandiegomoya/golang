package user

import "time"

//User is a new struc
type User struct {
	ID     int
	Name   string
	Date   time.Time
	Status bool
}

//altaUserimport new user
func (this *User) altaUser(id int, name string, date time.Time, status bool) {
	this.ID = id
	this.Name = name
	this.Date = date
	this.Status = status
}
