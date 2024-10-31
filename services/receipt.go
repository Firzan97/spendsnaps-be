package services

import (
	"log"
	"net/http"
	"spend-snap-be/models"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func Create(receipt *models.Receipt) *models.Receipt {

	if err := mgm.Coll(receipt).Create(receipt); err != nil {
		log.Print(err)
		panic(err)
	}

	return receipt
}

func GetAll() (receipts []*models.Receipt) {

	if err := mgm.Coll(&models.Receipt{}).SimpleFind(&receipts, bson.M{}); err != nil {
		log.Print(err)
		panic(err)
	}

	return
}

func GetOne(c *gin.Context) (receipt *models.Receipt) {
	id := c.Param("id")
	receipt = &models.Receipt{}

	err := mgm.Coll(receipt).FindByID(id, receipt)

	if err != nil {
		return nil
	}

	return
}

func Remove(c *gin.Context) (receipt *models.Receipt) {
	id := c.Param("id")
	receipt = &models.Receipt{}

	err := mgm.Coll(receipt).FindByID(id, receipt)
	if err != nil {
		return nil
	}

	// Delete this record
	if err := mgm.Coll(receipt).Delete(receipt); err != nil {
		panic(err)
	}

	return
}

func Update(c *gin.Context) (receipt *models.Receipt) {
	id := c.Param("id")

	var body models.UpdateReceipt
	receipt = &models.Receipt{}

	// Bind JSON payload to the Receipt struct
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := mgm.Coll(receipt).FindByID(id, receipt)
	if err != nil {
		return nil
	}

	if body.ShopName != nil {
		receipt.ShopName = *body.ShopName
	}
	if body.CompanyName != nil {
		receipt.CompanyName = *body.CompanyName
	}
	if body.Total != nil {
		receipt.Total = *body.Total
	}
	if body.Text != nil {
		receipt.Text = *body.Text
	}
	if body.ImageUrl != nil {
		receipt.ImageUrl = *body.ImageUrl
	}
	if body.Currency != nil {
		receipt.Currency = *body.Currency
	}

	// Update this record
	if err := mgm.Coll(receipt).Update(receipt); err != nil {
		panic(err)
	}

	return
}
