package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"fullName"`
	Email    string `json:"emailAddress"`
	Password string `json:"password"`
	Version  string `json:"__v"`
}

type JsonUser struct {
	Name     string `json:"fullName"`
	Email    string `json:"emailAddress"`
	Password string `json:"password"`
	Version  string `json:"__v"`
}