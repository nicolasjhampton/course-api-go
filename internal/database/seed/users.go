package seed

import (
	"fmt"
	"github.com/jinzhu/gorm"
	m "github.com/nicolasjhampton/course-api-go/internal/database/models"
)

var users = []m.User{
	m.User{
		ID:       3,
		Name:     "Joe Smith",
		Email:    "joe@smith.com",
		Password: "password",
	},
	m.User{
		ID:       2,
		Name:     "Sam Jones",
		Email:    "sam@jones.com",
		Password: "password",
	},
	m.User{
		ID:       1,
		Name:     "Sam Smith",
		Email:    "sam@smith.com",
		Password: "password",
	},
}

func userseeds(tx *gorm.DB) error {
	for _, user := range users {
		if err := tx.Create(&user).Error; err != nil {
			fmt.Println(err)
			tx.Rollback()
			return err
		}
	}
	return nil
}
