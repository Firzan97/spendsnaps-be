package routes

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {

	attach(r)
}

func attach(r *gin.Engine) {
	//Task
	Receipts(r)
}
