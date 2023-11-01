package handler

import (
	"github.com/gin-gonic/gin"
	"messenger-go/internal/logger"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logger.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
}
