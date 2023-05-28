package routes

import (
	"net/http"

	"github.com/SohamGhugare/shorten-url/database"
	"github.com/SohamGhugare/shorten-url/models"
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

	// Checking rate limit
	/*
		This function returns a gin.H object with error and rate limit reset
		if the rate limit is exceeded
	*/
	errH := database.CheckRateLimit(c.ClientIP())
	if errH != nil {
		c.JSON(http.StatusServiceUnavailable, errH)
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

	// Decrementing remaining requests
	rl := database.DecrRateLimit(c.ClientIP())

	c.JSON(http.StatusAccepted, gin.H{
		"body": models.Response{
			URL:             body.URL,
			CustomShort:     body.CustomShort,
			Expiry:          body.Expiry,
			XRateRemaining:  rl,
			XRateLimitReset: 30,
		},
	})
}
