package routes

import (
	"fmt"
	"net/http"

	"github.com/SohamGhugare/shorten-url/utility"
	"github.com/gin-gonic/gin"
)

func ShortenUrl(c *gin.Context) {

	body, err := utility.ValidateRequestBody(c)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Enforcing HTTP on the url
	body.URL = utility.EnforeHTTP(body.URL)

	c.JSON(http.StatusAccepted, gin.H{
		"body": body,
	})
}
