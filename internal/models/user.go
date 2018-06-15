package models

// import (
// 	"github.com/jinzhu/gorm"
// )

type User struct {
	ID			 int		`json:"_id"`
	Name     string `json:"fullName"`
	Email    string `json:"emailAddress"`
	Password string `json:"password"`
	Version  string `json:"__v"`
}