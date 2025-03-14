package utils

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJSON(c *gin.Context, payload any) error {
	if c.Request.Body == nil {
		return fmt.Errorf("missing request Body")
	}
	return json.NewDecoder(c.Request.Body).Decode(payload)
}

func WriteJSON(c *gin.Context, status int, v any) {
	c.IndentedJSON(status, v)
}

func WriteError(c *gin.Context, status int, err error) {
	c.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
}
