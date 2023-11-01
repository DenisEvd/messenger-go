package handler

import "github.com/gin-gonic/gin"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
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
