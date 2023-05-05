package register

import "github.com/gin-gonic/gin"

func UserRegister(e *gin.Engine) {
	e.POST("/v1/frontend/user/register", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "register success!",
		})
	})
}
