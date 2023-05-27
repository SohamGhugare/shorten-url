package main

import (
	"os"

	"github.com/SohamGhugare/shorten-url/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	// Running initializers
	initializers.LoadEnvVars()
}

func setupRoutes(r *gin.Engine) {
	// Setting up routes
}

func main(){
	r := gin.Default()

	setupRoutes(r)
	r.Run(os.Getenv("APP_PORT"))
}