package models


type Step struct {
	Id          uint `json:"_id"`
	CourseID    uint	 `json:"-"`
	Number      string `gorm:"AUTO_INCREMENT" json:"stepNumber"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
