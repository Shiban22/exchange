package router

import (
	"github.com/gin-gonic/gin"

	"exchange/controller"
	"exchange/mid"
)

func SetAuthRouter() *gin.Engine {

	r := gin.Default()

	auth := r.Group("/auth/api")
	{
		auth.POST("/login", controller.Login)
		//auth.POST("/register", func(ctx *gin.Context) {
		//	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		//		"message": "register,sucess",
		//	})
		//})
		auth.POST("/register", controller.Register)
	}
	api := r.Group("/api")
	api.Use(mid.AuthMiddleWare())
	{
		api.POST("/rate", controller.Rate)
	}

	return r
}
