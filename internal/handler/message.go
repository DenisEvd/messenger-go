package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"messenger-go/domain"
	"net/http"
	"strconv"
)

func (h *Handler) sendMessage(c *gin.Context) {
	var message domain.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid message json")
		return
	}

	fmt.Println(message)

	c.JSON(http.StatusOK, "")
}

func (h *Handler) getMessages(c *gin.Context) {
	senderID, err := strconv.Atoi(c.Param("sender_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid or not found sender id")
	}

	receiverID, err := strconv.Atoi(c.Param("sender_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid or not found receiver id")
	}

	fmt.Println(senderID, receiverID)

	c.JSON(http.StatusOK, "")
}
