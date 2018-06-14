package users

import (
	"github.com/gin-gonic/gin"
	m "github.com/nicolasjhampton/course-api-go/internal/models"
)

func GetUser(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(*m.User)
	c.JSON(200, user)
}

func GetUsers(c *gin.Context) {
	var users []m.User;
	DB.Find(&users)
	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
	var user m.User;
	if err := c.BindJSON(&user); err == nil {
		DB.FirstOrCreate(&user, user)
		c.JSON(200, user)
	} else {
		c.JSON(400, gin.H{})
	}	
}
