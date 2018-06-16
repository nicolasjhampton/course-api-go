package models

type Step struct {
	ID          uint   `json:"_id"`
	CourseID    uint   `json:"-"`
	Number      int    `json:"stepNumber"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
