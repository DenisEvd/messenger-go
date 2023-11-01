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

	api := router.Group("/chat")
	{
		api.POST("/", h.sendMessage)
		api.GET("/", h.getMessages)
	}

	return router
}
