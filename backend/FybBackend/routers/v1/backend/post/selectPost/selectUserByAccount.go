package selectPost

import (
	fybDatabase "FybBackend/database"
	"FybBackend/routers/v1/backend/token"
	"FybBackend/routers/v1/exceptionHandler"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func SelectPostByAccount(e *gin.Engine, db *gorm.DB) {
	e.GET("/v1/backend/post/searchByAccount", func(context *gin.Context) {
		if err := token.JwtVerify(context); err != nil {
			context.JSON(403, gin.H{
				"code":    403,
				"message": err.Error(),
			})
			return
		}
		var result *multierror.Error
		mp := make(map[string]interface{})
		account := context.DefaultQuery("account", "")
		if account == "" {
			result = multierror.Append(result, errors.New("账户输入不能为空！"))
		}
		mp["account"] = account
		post, _, err1 := fybDatabase.SelectSinglePostByCondition(db, mp)
		result = multierror.Append(result, err1)

		code, msg := exceptionHandler.Handle(result)
		if code == 200 {
			context.JSON(code, gin.H{
				"code":    code,
				"message": "get userInfoList success!",
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
