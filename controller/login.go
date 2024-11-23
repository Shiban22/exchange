package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"exchange/global"
	"exchange/models"
	"exchange/utils"
)

func Login(c *gin.Context) {
	var loginReq models.LoginReq
	var user models.User
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := global.DB.Where("username = ?", loginReq.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong credentials"})
		return
	}
	if !utils.CheckPassword(loginReq.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "pwd fild"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": "ok"})

}
