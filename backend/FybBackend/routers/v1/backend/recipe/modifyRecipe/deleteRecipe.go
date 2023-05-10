package modifyRecipe

import (
	fybDatabase "FybBackend/database"
	"FybBackend/routers/v1/backend/token"
	"FybBackend/routers/v1/exceptionHandler"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func DeleteRecipe(e *gin.Engine, db *gorm.DB) {
	e.DELETE("/v1/backend/recipe/delete", func(context *gin.Context) {
		if err := token.JwtVerify(context); err != nil {
			context.JSON(403, gin.H{
				"code":    403,
				"message": err.Error(),
			})
			return
		}

		var result *multierror.Error
		mp := make(map[string]interface{})
		mp["ID"] = context.DefaultQuery("id", "")
		_, err1 := fybDatabase.DeleteRecipe(db, mp)
		result = multierror.Append(result, err1)

		code, msg := exceptionHandler.Handle(result)
		if code == 200 {
			context.JSON(code, gin.H{
				"code":    code,
				"message": "删除成功",
			})
		} else {
			context.JSON(code, gin.H{
				"code":    code,
				"message": msg,
			})
		}
	})
}
