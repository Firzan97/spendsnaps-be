package main

import (
	"os"
	"spend-snap-be/routes"

	"spend-snap-be/config"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load config
	if err := config.Init(); err != nil {
		return
	}

	// Connect to Database
	if err := config.App.Db.Connect(); err != nil {
		return
	}

	r := gin.Default()

	routes.Init((r))

	r.Run(":" + os.Getenv("PORT"))
}
