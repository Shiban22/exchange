package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"exchange/global"
	"exchange/models"
	"exchange/utils"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	HashPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Password = HashPwd
	if err := global.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	token, err := utils.GenerateJwt(user.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"CreateTokenerror": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
