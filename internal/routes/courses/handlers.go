package courses

import (
	"fmt"
	"github.com/gin-gonic/gin"
	m "github.com/nicolasjhampton/course-api-go/internal/models"
	"net/http"
)

func GetCourses(c *gin.Context) {
	var courses []m.Course
	var jArr []interface{}
	DB.Table("Courses").Select("ID, Title").Find(&courses)
	for _, course := range courses {
		jArr = append(jArr, struct {
			ID    uint   `json:"_id"`
			Title string `json:"title"`
		}{course.ID, course.Title})
	}
	c.JSON(http.StatusOK, jArr)
}

func FindCourse(c *gin.Context) {
	var course m.Course
	id := c.Param("courseid")
	DB.Preload("User").Preload("Reviews").First(&course, id)
	c.Set("course", course)
}

func GetCourse(c *gin.Context) {
	course := c.MustGet("course").(m.Course)
	c.JSON(http.StatusOK, course)
}

func CreateCourse(c *gin.Context) {
	var course m.Course
	var err error
	if err = c.BindJSON(&course); err == nil {
		user := c.MustGet(gin.AuthUserKey).(*m.User)
		DB.Where(m.Course{Title: course.Title}).Assign(m.Course{UserID: user.ID}).FirstOrCreate(&course)
		c.Redirect(http.StatusCreated, fmt.Sprintf("/api/v1/courses/%v", course.ID))
	} else {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
	}
}
