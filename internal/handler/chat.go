package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type createChatInput struct {
	Name string `json:"name" binding:"required"`
}

type addUserInput struct {
	UserID int `json:"user_id" binding:"required"`
	ChatID int `json:"chat_id" binding:"required"`
}

func (h *Handler) createChat(c *gin.Context) {
	var input createChatInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid chat name")
		return
	}

	id, err := h.services.Chat.Create(input.Name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getUserChats(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid sender id")
		return
	}

	chats, err := h.services.GetUserChats(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, chats)
}

func (h *Handler) addUserToChat(c *gin.Context) {
	var input addUserInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := h.services.Chat.AddUser(input.ChatID, input.UserID); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}
