package services

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"spend-snap-be/models"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/textract"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create(receipt *models.Receipt) *models.Receipt {

	if err := mgm.Coll(receipt).Create(receipt); err != nil {
		log.Print(err)
		panic(err)
	}

	return receipt
}

func GetAll() (receipts []*models.Receipt) {
	findOptions := options.Find().SetSkip(int64(0)).SetLimit(int64(30))

	if err := mgm.Coll(&models.Receipt{}).SimpleFind(&receipts, bson.M{}, findOptions); err != nil {
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

func Extract(c *gin.Context) {
	// Parse the multipart form to get the file
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil { // Limit upload size to 10 MB
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse form data"})
		return
	}

	// Get the file from the form data
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not get file from form data"})
		return
	}

	// Open the file
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not open file"})
		return
	}
	defer file.Close()

	// Read the file into a byte slice
	imageData, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read file"})
		return
	}

	// Upload the image to S3
	s3Key := fmt.Sprintf("receipts/%d.jpg", time.Now().UnixNano()) // Unique key based on timestamp
	if err := uploadToS3(imageData, s3Key); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload to S3: " + err.Error()})
		return
	}

	// Call Textract to analyze the image in S3
	extractedData, err := analyzeReceipt(imageData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Textract error: " + err.Error()})
		return
	}

	// Return success message and the extracted data
	c.JSON(http.StatusOK, gin.H{"message": "Receipt processed successfully", "data": extractedData})

}

func analyzeReceipt(imageData []byte) (string, error) {
	// Load AWS config
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", fmt.Errorf("failed to load config, %v", err)
	}

	textractClient := textract.NewFromConfig(cfg)

	// Prepare the input for Textract
	input := &textract.DetectDocumentTextInput{
		Document: &types.Document{
			Bytes: imageData,
		},
	}

	// Call Textract to analyze the document
	result, err := textractClient.DetectDocumentText(context.TODO(), input)
	if err != nil {
		return "", fmt.Errorf("failed to analyze document: %v", err)
	}

	// Extract only the text blocks
	var extractedText string
	for _, block := range result.Blocks {
		if block.BlockType == types.BlockTypeLine {
			extractedText += *block.Text + "\n" // Append each line of text
		}
	}

	return extractedText, nil
}

func uploadToS3(data []byte, key string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return fmt.Errorf("failed to load config, %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)

	// Upload the data to S3
	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String("spend-snap"),
		Key:         aws.String(key),
		Body:        bytes.NewReader(data),
		ContentType: aws.String("image/jpeg"),
	})

	return err
}
