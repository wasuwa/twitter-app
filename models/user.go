package models

import (
	"strconv"
	"twitter-app/database"
)

type User struct {
	Id       int        `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
}

func (u *User) All() []User {
	var users []User
	d := database.GetDB()
	d.Find(&users)
	return users
}

func (u *User) Create() {
	d := database.GetDB()
	d.Create(u)
}

func (u *User) Find(id int) {
	u.Id = id
	d := database.GetDB()
	d.First(u)
}

func ConvertToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return i, err
	}
	return i, err
}
