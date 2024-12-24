package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"exchange/global"
	"exchange/models"
)

func Rate(c *gin.Context) {
	var rate models.Rate
	if err := c.ShouldBindJSON(&rate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rate.Date = time.Now()
	global.DB.AutoMigrate(rate)
	global.DB.Create(rate)
	c.JSON(http.StatusOK, rate)
}

func getrate(c *gin.Context) {
	var rates []models.Rate
	if err := global.DB.Find(&rates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Rate)
}
