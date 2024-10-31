package routes

import (
	"spend-snap-be/controllers"

	"github.com/gin-gonic/gin"
)

func Receipts(r *gin.Engine) {
	api := r.Group("/receipts")
	{

		api.POST("/extract", controllers.ExtractReceipt)
		api.GET("/", controllers.GetAll)
		api.GET("/:id", controllers.GetOne)
		api.POST("/", controllers.Create)
		api.PUT("/:id", controllers.Update)
		api.DELETE("/:id", controllers.Remove)
	}
}
