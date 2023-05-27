package utility

import (
	"errors"
	"fmt"

	"github.com/SohamGhugare/shorten-url/models"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

/*
VALIDATE REQUEST BODY
This function checks for any defects in the request body and binds it to the model
Checks:

	url: not null, valid
	short: not null
*/
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

	// Checking for invalid url
	if !govalidator.IsURL(body.URL) {
		return nil, errors.New("invalid url")
	}

	// Checking for empty short
	if body.CustomShort == "" {
		return nil, errors.New("short cannot be empty")
	}
	return body, nil
}
