package handler

import (
	"MoneyMinder/internal/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetProfits(c *gin.Context) {
	profits, err := h.services.Records.GetProfitRecords()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "profits not founded"})
		return
	}
	if len(profits) == 0 {
		c.JSON(http.StatusOK, gin.H{"profits": []*dtos.RecordOutput{}})
		return
	}

	c.JSON(200, gin.H{"profits": profits})
}
func (h *Handler) GetExpences(c *gin.Context) {
	expenses, err := h.services.Records.GetExpenseRecords()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "expenses not founded"})
		return
	}
	if len(expenses) == 0 {
		c.JSON(http.StatusOK, gin.H{"expenses": []*dtos.RecordOutput{}})
		return
	}

	c.JSON(200, gin.H{"expenses": expenses})
}
func (h *Handler) PostRecord(c *gin.Context) {
	var request dtos.RecordInput

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"message": "Invalid body request"})
		return
	}

	if err := h.services.Records.AddRecord(request); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "The record has been added"})
}

func (h *Handler) PutRecord(c *gin.Context) {
	recordId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid id"})
		return
	}

	var request dtos.RecordInput
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"message": "Invalid body request"})
		return
	}

	if err := h.services.Records.UpdateRecord(recordId, request); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "The record has been updated"})
}
