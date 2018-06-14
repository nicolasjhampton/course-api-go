package courses

import (
	"github.com/gin-gonic/gin"
)

func GetCourses(c *gin.Context) {
	var courses []Course;
	DB.Find(&courses)
	c.JSON(200, courses)
}

func GetCourse(c *gin.Context) {
	var course Course;
	id := c.Param("id")
	DB.First(&course, id)
	c.JSON(200, course)
}
