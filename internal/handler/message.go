package handler

import (
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

	id, err := h.services.Message.Create(message)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "sending message error")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getMessages(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid sender id")
		return
	}

	chatID, err := strconv.Atoi(c.Query("chat_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid receiver id")
		return
	}

	messages, err := h.services.Message.GetAll(userID, chatID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "error with getting chat messages")
	}

	c.JSON(http.StatusOK, messages)
}
