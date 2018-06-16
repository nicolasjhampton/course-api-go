package users

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	auth "github.com/nicolasjhampton/course-api-go/internal/middleware/authorization"
)

var DB *gorm.DB

func Routes(g gin.IRouter, db *gorm.DB) *gin.RouterGroup {
	DB = db
	seedUsers()
	users := g.Group("/users")
	{
		users.GET("/", auth.Required(DB), GetUser)
		users.POST("/", CreateUser)
		users.GET("/all", auth.Admin(DB), GetUsers)
	}
	user := users.Group("/:id")
	return user
}
