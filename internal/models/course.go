package models

import (
    "github.com/jinzhu/gorm"
)

type Course struct {
	gorm.Model
	//User        User     `json:"user"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Time        string   `json:"estimatedTime"`
	Materials   string   `json:"materialsNeeded"`
	Version     string   `json:"__v"`
	//Reviews     []Review `json:"reviews"`
	//Steps       []Step   `json:"steps"`
}

// type Review struct {
// 	Id      string `json:"_id"`
// 	UserId  string `json:"user"`
// 	Rating  string `json:"rating"`
// 	Review  string `json:"review"`
// 	Version string `json:"__v"`
// }

// type Course struct {
//     //gorm.Model
// 	ID    string //`json:"_id" gorm:"primary_key"`
// 	Title string `json:"title"`
// }

// type Step struct {
// 	Id          string `json:"_id"`
// 	Number      string `json:"stepNumber"`
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// }
