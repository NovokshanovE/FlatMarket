package handlers

import (
	"net/http"

	"github.com/NovokshanovE/flatmarket/database"
	"github.com/NovokshanovE/flatmarket/models"
	"github.com/gin-gonic/gin"
)

func CreateHouse(c *gin.Context) {
	var house models.House
	if err := c.ShouldBindJSON(&house); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&house).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, house)
}

func GetFlats(c *gin.Context) {
	houseID := c.Param("id")
	var flats []models.Flat
	userType := c.GetString("userType")

	if userType == "client" {
		database.DB.Where("house_id = ? AND status = ?", houseID, "approved").Find(&flats)
	} else {
		database.DB.Where("house_id = ?", houseID).Find(&flats)
	}

	c.JSON(http.StatusOK, flats)
}

func SubscribeHouse(c *gin.Context) {
	// Implementation of subscription logic
	c.JSON(http.StatusOK, gin.H{"message": "Subscribed successfully"})
}
