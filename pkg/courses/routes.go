package courses

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB;

func Routes(g gin.IRouter, db *gorm.DB) {
	DB = db;
	seedCourses()
	courses := g.Group("/courses")
	{
		courses.GET("/", GetCourses)
		courses.GET("/:id", GetCourse)
	}
}