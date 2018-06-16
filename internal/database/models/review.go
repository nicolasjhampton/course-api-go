package models

import (
	"github.com/jinzhu/gorm"
)

type Review struct {
	gorm.Model
	UserID  uint `json:"user"`
	CourseID  uint `json:"-"`
	Rating  string `json:"rating"`
	Review  string `json:"review"`
	Version string `json:"__v"`
}
