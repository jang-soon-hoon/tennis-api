package routes

import (
	"tennis-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	api := r.Group("/api")
	{
		// Matches route
		api.GET("/matches", controllers.GetMatches)
	}
}
