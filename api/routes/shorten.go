package routes

import (
	"net/http"

	"github.com/SohamGhugare/shorten-url/database"
	"github.com/SohamGhugare/shorten-url/utility"
	"github.com/gin-gonic/gin"
)

func ShortenUrl(c *gin.Context) {

	body, err := utility.ValidateRequestBody(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Enforcing HTTP on the url
	body.URL = utility.EnforeHTTP(body.URL)

	// Storing the url in database
	if err := database.CreateURL(body.URL, body.CustomShort); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"body": body,
	})
}
