package utility

import (
	"errors"
	"fmt"

	"github.com/SohamGhugare/shorten-url/models"
	"github.com/gin-gonic/gin"
)

func ValidateRequestBody(c *gin.Context) (*models.Request, error) {
	body := new(models.Request)

	// Binding the request body with the model
	if err := c.Bind(&body); err != nil {
		return nil, fmt.Errorf("error parsing request body: %v", err)
	}

	// Checking for empty url
	if body.URL == "" {
		return nil, errors.New("url cannot be empty")
	}

	// Checking for empty short
	if body.CustomShort == "" {
		return nil, errors.New("short cannot be empty")
	}
	return body, nil
}
