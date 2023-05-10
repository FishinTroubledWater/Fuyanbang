package modifyPost

import (
	fybDatabase "FybBackend/database"
	"FybBackend/routers/v1/backend/token"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func DeletePost(e *gin.Engine, db *gorm.DB) {
	e.DELETE("/v1/backend/post/delete", func(context *gin.Context) {
		var result *multierror.Error
		mp := make(map[string]interface{})
		err1 := token.JwtVerify(context)

		mp["id"] = context.DefaultQuery("id", "")
		_, err2 := fybDatabase.DeletePost(db, mp)
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
