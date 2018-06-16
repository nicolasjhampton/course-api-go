package DB

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	m "github.com/nicolasjhampton/course-api-go/internal/database/models"
	"github.com/nicolasjhampton/course-api-go/internal/database/seed"
	"errors"
)

func Setup() (*gorm.DB, error) {
	// connect
	DB, err := gorm.Open("sqlite3", "test.db"); 
	if err != nil {
		return nil, errors.New("failed to connect database")
	}
	
	// Create tables
	DB.DropTable("Courses", "Steps", "Reviews", "Users")
	DB.AutoMigrate(&m.User{}, &m.Review{}, &m.Step{}, &m.Course{})

	// Seed database
	if err = seed.RunTx(DB); err != nil {
		return nil, err
	}
	
	return DB, nil
}