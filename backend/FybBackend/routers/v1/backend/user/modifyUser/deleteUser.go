package modifyUser

import (
	fybDatabase "FybBackend/database"
	"FybBackend/routers/v1/backend/token"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func DeleteUser(e *gin.Engine, db *gorm.DB) {
	e.DELETE("/v1/backend/user/delete", func(context *gin.Context) {
		var result *multierror.Error
		mp := make(map[string]interface{})
		err1 := token.JwtVerify(context)

		mp["account"] = context.DefaultQuery("account", "")
		_, err2 := fybDatabase.DeleteUser(db, mp)
		result = multierror.Append(result, err1, err2)

		code := 200
		if result.ErrorOrNil() == nil {
			context.JSON(code, gin.H{
				"code":    code,
				"message": "删除成功",
			})
		} else {
			if err1 != nil {
				code = 403
			} else if err2 != nil {
				code = 404
			} else {
				code = 500
			}
			context.JSON(code, gin.H{
				"code":    code,
				"message": result.Error(),
			})

		}
	})
}
