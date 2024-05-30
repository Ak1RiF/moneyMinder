package handler

import (
	"MoneyMinder/internal/dtos"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) GetGoals(c *gin.Context) {
	goals, err := h.services.Goals.GetAllGoals()
	if err != nil {
		c.JSON(404, gin.H{"message": err.Error()})
		return
	}
	if len(goals) == 0 {
		c.JSON(200, gin.H{"goals": []*dtos.GoalOutput{}})
		return
	}

	c.JSON(200, gin.H{"goals": goals})
}

func (h *Handler) GetGoal(c *gin.Context) {
	goalId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid id"})
		return
	}

	goal, err := h.services.Goals.GetGoalById(goalId)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"goal": goal})
}

func (h *Handler) CreateGoal(c *gin.Context) {
	var request dtos.CreateGoal

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"message": "Invalid body request"})
		return
	}

	if err := h.services.Goals.CreateNewGoal(request); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "The goal has been created"})
}

func (h *Handler) UpdateGoal(c *gin.Context) {
	goalId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid id"})
		return
	}

	var request dtos.UpdateGoal

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"message": "Invalid body request"})
		return
	}
	if err := h.services.Goals.UpdateFieldGoal(goalId, request); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "The goal has been updated"})
}
func (h *Handler) RemoveGoal(c *gin.Context) {
	goalId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid id"})
		return
	}
	if err := h.services.Goals.RemoveGoalById(goalId); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "The goal has been deleted"})
}
