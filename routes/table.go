package routes

import (
	"github.com/gin-gonic/gin"
	"go-code-generation/controllers"
)

func apiRoute(r *gin.Engine) {
	api := r.Group("/tables")

	api.GET("/", controllers.GetTables)

	api.GET("/:id", controllers.GetTable)
}
