package reviews

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	auth "github.com/nicolasjhampton/course-api-go/internal/middleware/authorization"
)

var DB *gorm.DB

func Routes(g gin.IRouter, db *gorm.DB) *gin.RouterGroup {
	DB = db
	reviews := g.Group("/reviews")
	{
		reviews.POST("/", auth.Required(DB), CreateReview)
	}
	review := reviews.Group("/:reviewid")
	return review
}
