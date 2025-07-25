package controllers

import (
	"net/http"
	"time"

	"activity-tracker/config"
	"activity-tracker/models"
	"github.com/gin-gonic/gin"
)

type LogHoursRequest struct {
	ActivityID uint    `json:"activity_id" binding:"required"`
	Hours      float64 `json:"hours" binding:"required"`
}

// Зафиксировать потраченные часы
func LogActivityHours(c *gin.Context) {
	var request LogHoursRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверяем существование активности
	var activity models.Activity
	if err := config.DB.First(&activity, request.ActivityID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Активность не найдена"})
		return
	}

	// Создаем запись лога
	activityLog := models.ActivityLog{
		ActivityID: request.ActivityID,
		Hours:      request.Hours,
		LoggedAt:   time.Now(),
	}

	config.DB.Create(&activityLog)
	c.JSON(http.StatusCreated, gin.H{"log": activityLog})
}
