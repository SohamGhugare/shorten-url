package routes

import (
	"net/http"

	"github.com/SohamGhugare/shorten-url/database"
	"github.com/gin-gonic/gin"
)

// Resolving the url
func ResolveUrl(c *gin.Context) {
	short := c.Params.ByName("url")
	url, err := database.GetURL(short)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Redirect(http.StatusPermanentRedirect, url)
	
}
