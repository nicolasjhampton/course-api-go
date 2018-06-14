package courses

import (
	"github.com/gin-gonic/gin"
	m "github.com/nicolasjhampton/course-api-go/internal/models"
)

func GetCourses(c *gin.Context) {
	var courses []m.Course;
	DB.Find(&courses)
	c.JSON(200, courses)
}

func GetCourse(c *gin.Context) {
	var course m.Course;
	id := c.Param("id")
	DB.First(&course, id)
	c.JSON(200, course)
}
