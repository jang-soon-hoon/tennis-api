package routes

import (
	"tennis-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Matches route
	r.GET("/matches", controllers.GetMatches)
}
