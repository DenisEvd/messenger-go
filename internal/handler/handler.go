package handler

import (
	"github.com/gin-gonic/gin"
	"messenger-go/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	chat := router.Group("/chat")
	{
		chat.POST("/", h.createChat)

		user := chat.Group("/user")
		{
			user.POST("/", h.addUserToChat)
			user.GET("/", h.getUserChats)
		}
	}

	msg := router.Group("/message")
	{
		msg.POST("/", h.sendMessage)
		msg.GET("/", h.getMessages)
	}

	return router
}
