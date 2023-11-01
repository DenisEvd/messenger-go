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
	senderID, err := strconv.Atoi(c.Query("sender_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid sender id")
		return
	}

	receiverID, err := strconv.Atoi(c.Query("receiver_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid receiver id")
		return
	}

	fmt.Println(senderID, receiverID)

	c.JSON(http.StatusOK, "")
}
