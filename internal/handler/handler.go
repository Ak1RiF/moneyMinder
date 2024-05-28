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
			records.PUT("/:id", h.PutRecord)
		}

		goals := api.Group("/goals")
		{
			goals.GET("/", h.GetGoals)
			goals.GET("/:id", h.GetGoal)
			goals.POST("/", h.CreateGoal)
			goals.PUT("/:id", h.UpdateGoal)
			goals.PUT("/:id", h.UpdateGoal)
			goals.DELETE("/:id", h.RemoveGoal)
		}
	}

	return router
}
