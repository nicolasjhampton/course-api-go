package users

import (
	m "github.com/nicolasjhampton/course-api-go/internal/models"
)

func seedUsers() {
	DB.DropTable(&m.User{})
	DB.AutoMigrate(&m.User{})
	DB.Create(&m.User{
		Name: "Joe Smith",
		Email: "joe@smith.com",
		Password: "password",
	})
	DB.Create(&m.User{
		Name: "Sam Jones",
		Email: "sam@jones.com",
		Password: "password",
	})
	DB.Create(&m.User{
		Name: "Sam Smith",
		Email: "sam@smith.com",
		Password: "password",
	})
}