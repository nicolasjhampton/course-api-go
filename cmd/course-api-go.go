package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nicolasjhampton/course-api-go/internal/database/DB"
	"github.com/nicolasjhampton/course-api-go/internal/routes/courses"
	"github.com/nicolasjhampton/course-api-go/internal/routes/reviews"
	"github.com/nicolasjhampton/course-api-go/internal/routes/users"
)

func main() {
	var db *gorm.DB
	var err error
	if db, err = DB.Setup(); err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := gin.Default()

	api := router.Group("/api")

	v1 := api.Group("/v1")

	_ = users.Routes(v1, db)

	course := courses.Routes(v1, db)

	_ = reviews.Routes(course, db)

	router.Run(os.Getenv("PORT"))
}
