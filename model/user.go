package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Id        uint64 `json:"id" form:"-"`
	Username  string `json:"username" form:"username"`
	Password  string `json:"-" form:"password"`
	FirstName string `json:"firstname" form:"firstname"`
	LastName  string `json:"lastname" form:"lastname"`
	Email     string `json:"email" form:"email"`
	Created   string `json:"created" form:"created"`
}

type Users []User