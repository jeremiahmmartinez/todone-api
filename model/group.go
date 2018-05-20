package model

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	Id        string `json:"id" form:"-"`
	OwnerId   string `json:"owner_id" form:"owner_id"`
	Label	  string `json:"label" form:"label"`
	Description string `json:"description" form:"description"`
	Created   string `json:"created" form:"created"`
	Updated   string `json:"updated" form:"updated"`
}

type Groups []Group