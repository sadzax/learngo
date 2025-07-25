package controllers

import (
	"activity-tracker/config"
	"activity-tracker/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// C
func CreateActivity(c *gin.Context) {
	var activity models.Activity

	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&activity)
	c.JSON(http.StatusCreated, gin.H{"activity": activity})
}

// R
func GetUserActivities(c *gin.Context) {
	userID := c.Param("user_id")
	var activities []models.Activity

	config.DB.Where("user_id = ?", userID).Find(&activities)
	c.JSON(http.StatusOK, gin.H{"activities": activities})
}

// U
func UpdateActivity(c *gin.Context) {
	id := c.Param("id")
	var activity models.Activity

	if err := config.DB.First(&activity, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Активность не найдена"})
		return
	}

	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&activity)
	c.JSON(http.StatusOK, gin.H{"activity": activity})
}

// D
func DeleteActivity(c *gin.Context) {
	id := c.Param("id")
	var activity models.Activity

	if err := config.DB.First(&activity, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Активность не найдена"})
		return
	}

	config.DB.Delete(&activity)
	c.JSON(http.StatusOK, gin.H{"message": "Активность удалена"})
}
