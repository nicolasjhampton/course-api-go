package reviews

import (
	m "github.com/nicolasjhampton/course-api-go/internal/models"
)

func seedReviews() {
	DB.DropTable(&m.Review{})
	DB.AutoMigrate(&m.Review{})
	DB.Create(&m.Review{
		UserID: 1,
		CourseID: 1,
		Rating: "5",
		Review: "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
		Version: "0",
	})
	DB.Create(&m.Review{
		UserID: 2,
		CourseID: 1,
		Rating: "3",
		Review: "Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		Version: "0",
	})
	DB.Create(&m.Review{
		UserID: 1,
		CourseID: 1,
		Rating: "5",
		Review: "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		Version: "0",
	})
}