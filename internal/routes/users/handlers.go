package users

import (
	"github.com/gin-gonic/gin"
	m "github.com/nicolasjhampton/course-api-go/internal/models"
)

func GetUser(c *gin.Context) {
	var user m.User;
	DB.First(&user/*, id*/)
	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	var user m.User;
	var juser m.JsonUser;
	if err := c.BindJSON(&juser); err == nil {
		DB.FirstOrCreate(&user, m.User{
			Name: juser.Name,
			Email: juser.Email,
			Password: juser.Password,
			Version: "0",
		})
		c.JSON(200, user)
	} else {
		c.JSON(400, gin.H{})
	}	
}
