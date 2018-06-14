package users

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "github.com/nicolasjhampton/course-api-go/internal/models"
)

var DB *gorm.DB;

func Routes(g gin.IRouter, db *gorm.DB) {
	DB = db;
	DB.DropTable(&m.User{})
	DB.AutoMigrate(&m.User{})
	//seedUsers()
	users := g.Group("/users")
	{
		users.GET("/", GetUser)
		users.POST("/", CreateUser)
	}
}