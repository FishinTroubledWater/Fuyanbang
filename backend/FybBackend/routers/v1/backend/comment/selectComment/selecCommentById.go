package selectComment

import (
	fybDatabase "FybBackend/database"
	"FybBackend/routers/v1/backend/token"
	"FybBackend/routers/v1/exceptionHandler"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func SelectCommentById(e *gin.Engine, db *gorm.DB) {
	e.GET("/v1/backend/comment/searchById", func(context *gin.Context) {
		if err := token.JwtVerify(context); err != nil {
			context.JSON(403, gin.H{
				"code":    403,
				"message": err.Error(),
			})
			return
		}

		var result *multierror.Error
		mp := make(map[string]interface{})
		mp["comment.id"] = context.DefaultQuery("id", "")
		post, _, err1 := fybDatabase.SelectSingleCommentByCondition(db, mp)
		result = multierror.Append(result, err1)

		code, msg := exceptionHandler.Handle(result)
		if code == 200 {
			context.JSON(code, gin.H{
				"code":    code,
				"message": "请求成功",
				"data":    post,
			})
		} else {
			context.JSON(code, gin.H{
				"code":    code,
				"message": msg,
			})
		}
	})
}
