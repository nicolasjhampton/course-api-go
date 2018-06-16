package models

type Review struct {
	ID          uint `json:"_id"`
	UserID  uint `json:"user"`
	CourseID  uint `json:"-"`
	Rating  string `json:"rating"`
	Review  string `json:"review"`
	Version string `json:"__v"`
}
