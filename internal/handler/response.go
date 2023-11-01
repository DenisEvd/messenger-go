package handler

import "github.com/gin-gonic/gin"

func newErrorResponse(c *gin.Context, status int) {
	c.AbortWithStatus(status)
}
