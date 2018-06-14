package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
    "github.com/nicolasjhampton/course-api-go/internal/routes/courses"
    "github.com/nicolasjhampton/course-api-go/internal/routes/users"
)

func main() {

    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

	router := gin.Default()

    api := router.Group("/api")
    
    v1 := api.Group("/v1")

    courses.Routes(v1, db)
    
    users.Routes(v1, db)

    
    
	router.Run(":5000") // listen and serve on 0.0.0.0:8080
}
