package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"messenger-go/domain"
	"net/http"
)

func (h *Handler) sendMessage(c *gin.Context) {
	var message domain.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		newErrorResponse(c, http.StatusBadRequest)
		return
	}

	fmt.Println(message)

	c.JSON(http.StatusOK, "")
}

func (h *Handler) getMessages(c *gin.Context) {

}
