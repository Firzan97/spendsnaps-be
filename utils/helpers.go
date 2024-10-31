package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetDefaultString will set value to backup if it's
// an empty string
func SetDefaultString(value *string, backup string) {
	if len(*value) == 0 {
		*value = backup
	}
}

func SendMessage(c *gin.Context, code int, message string) {
	// Prepare error code for response body
	err := code
	// Set to 0 if code is 200
	if code == http.StatusOK {
		err = 0
	}
	// Abort & send response back
	c.AbortWithStatusJSON(code, gin.H{"error": err, "message": message})
}

// InternalError - HTTP 500 Error
func InternalError(c *gin.Context, message string) {
	SetDefaultString(&message, "An Internal Error occurred.")
	SendMessage(c, http.StatusInternalServerError, message)
}

// ServiceUnavailable - HTTP 502 Error
func ServiceUnavailable(c *gin.Context, message string) {
	SetDefaultString(&message, "Service unavailable.")
	SendMessage(c, http.StatusServiceUnavailable, message)
}

// Success - HTTP 200 OK
func Success(c *gin.Context, message string) {
	SetDefaultString(&message, "Success.")
	SendMessage(c, http.StatusOK, message)
}

// NotFound - HTTP 404 Not Found
func NotFound(c *gin.Context, message string) {
	SetDefaultString(&message, "Not Found.")
	SendMessage(c, http.StatusNotFound, message)
}

// Unauthorized - HTTP 401 Unauthorized
func Unauthorized(c *gin.Context, message string) {
	SetDefaultString(&message, "Unauthorized.")
	SendMessage(c, http.StatusUnauthorized, message)
}

// Forbidden - HTTP 403 Forbidden
func Forbidden(c *gin.Context, message string) {
	SetDefaultString(&message, "Forbidden.")
	SendMessage(c, http.StatusForbidden, message)
}

// BadRequest - HTTP 400 Bad Request
func BadRequest(c *gin.Context, message string) {
	SetDefaultString(&message, "Bad Request.")
	SendMessage(c, http.StatusBadRequest, message)
}

// SuccessData - send Data
func SuccessData(c *gin.Context, message string, data interface{}) {
	SetDefaultString(&message, "Success.")
	c.AbortWithStatusJSON(
		http.StatusOK,
		gin.H{
			"error":   0,
			"message": message,
			"data":    data,
		},
	)
}

// Status of API
func Status(c *gin.Context) {
	Success(c, "API is running...")
}

// OnNotFound - 404
func OnNotFound(c *gin.Context) {
	NotFound(c, "")
}

// MissingInfo - HTTP 400 Bad Request
func MissingInfo(c *gin.Context, message string) {
	SetDefaultString(&message, "Invalid or missing information.")
	SendMessage(c, http.StatusBadRequest, message)
}
