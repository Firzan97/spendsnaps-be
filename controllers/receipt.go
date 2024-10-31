package controllers

import (
	"net/http"
	"spend-snap-be/models"
	"spend-snap-be/services"

	"github.com/gin-gonic/gin"

	"spend-snap-be/utils"
)

func ExtractReceipt(c *gin.Context) {

}

func GetAll(c *gin.Context) {
	docs := services.GetAll()

	utils.SuccessData(c, "", docs)
}

func Create(c *gin.Context) {
	var receipt models.Receipt

	// Bind JSON payload to the Receipt struct
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdReceipt := services.Create(&receipt)

	utils.SuccessData(c, "Receipt successfully created", createdReceipt)
}

func Update(c *gin.Context) {
	doc := services.Update(c)

	if doc == nil {
		utils.NotFound(c, "Receipt not found")
		return
	}

	utils.SuccessData(c, "Receipt successfully updated", doc)
}

func Remove(c *gin.Context) {
	doc := services.Remove(c)

	if doc == nil {
		utils.NotFound(c, "Receipt not found")
		return
	}

	utils.Success(c, "Receipt successfully deleted")

}

func GetOne(c *gin.Context) {
	doc := services.GetOne(c)

	if doc == nil {
		utils.NotFound(c, "Receipt not found")
		return
	}

	utils.SuccessData(c, "", doc)
}
