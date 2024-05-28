package handler

import (
	"MoneyMinder/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		records := api.Group("/records")
		{
			records.GET("/profits", h.GetProfits)
			records.GET("/expenses", h.GetExpences)
			records.POST("/", h.PostRecord)
			records.PUT("/", h.PutRecord)
		}
	}

	return router
}
