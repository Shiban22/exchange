package router

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"exchange/controller"
)

func SetAuthRouter() *gin.Engine {

	r := gin.Default()

	auth := r.Group("/auth/api")
	{
		auth.POST("/login", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"message": "login,sucess",
			})
		})
		//auth.POST("/register", func(ctx *gin.Context) {
		//	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		//		"message": "register,sucess",
		//	})
		//})
		auth.POST("/register", controller.Register)
	}
	return r
}

func register(context *gin.Context) {

}
