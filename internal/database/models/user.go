package models

type User struct {
	ID       uint    `json:"_id"`
	Name     string `json:"fullName"`
	Email    string `json:"emailAddress"`
	Password string `json:"password"`
	Version  string `json:"__v"`
}
