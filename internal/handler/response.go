package handler

import "github.com/gin-gonic/gin"

type Error struct {
	Message string `json:"Message"`
}

func NewErrorResponse(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, Error{Message: message})
	return
}
