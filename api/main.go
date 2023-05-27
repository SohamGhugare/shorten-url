package main

import (
	"os"

	"github.com/SohamGhugare/shorten-url/database"
	"github.com/SohamGhugare/shorten-url/initializers"
	"github.com/SohamGhugare/shorten-url/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	// Running initializers
	initializers.LoadEnvVars()
	database.CreateClient(0)
}

func setupRoutes(r *gin.Engine) {
	// Setting up routes
	r.POST("/api/v1", routes.ShortenUrl)
	r.GET("/:url", routes.ResolveUrl)
}

func main() {
	r := gin.Default()

	setupRoutes(r)
	r.Run(os.Getenv("APP_PORT"))
}
