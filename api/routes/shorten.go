package routes

import (
	"net/http"

	"github.com/SohamGhugare/shorten-url/models"
	"github.com/gin-gonic/gin"
)

func ShortenUrl(c *gin.Context) {
	body := new(models.Request)
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error parsing request body",
		})
	}
	c.JSON(http.StatusAccepted, gin.H{
		"body": body,
	})
}
