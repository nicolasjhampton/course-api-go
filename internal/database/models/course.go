package models

type Course struct {
	ID          uint `json:"_id"`
	UserID      uint     `json:"-"`
	User        User     `json:"user"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Time        string   `json:"estimatedTime"`
	Materials   string   `json:"materialsNeeded"`
	Version     string   `json:"__v"`
	Reviews     []Review `json:"reviews"`
	Steps       []Step   `json:"steps"`
}