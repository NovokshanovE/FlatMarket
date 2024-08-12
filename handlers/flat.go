package handlers

import (
	"net/http"

	"github.com/NovokshanovE/flatmarket/database"
	"github.com/NovokshanovE/flatmarket/models"
	"github.com/gin-gonic/gin"
)

func CreateFlat(c *gin.Context) {
	var flat models.Flat
	if err := c.ShouldBindJSON(&flat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	flat.Status = "created"

	if err := database.DB.Create(&flat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var house models.House
	if err := database.DB.First(&house, flat.HouseID).Error; err == nil {
		house.LastFlatAddedAt = flat.CreatedAt
		database.DB.Save(&house)
	}

	c.JSON(http.StatusOK, flat)
}

func UpdateFlat(c *gin.Context) {
	var req struct {
		ID     uint   `json:"id"`
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var flat models.Flat
	if err := database.DB.First(&flat, req.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flat not found"})
		return
	}

	flat.Status = req.Status

	if err := database.DB.Save(&flat).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, flat)
}
