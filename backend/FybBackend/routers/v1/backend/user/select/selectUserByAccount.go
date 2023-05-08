package _select

import (
	fybDatabase "FybBackend/database"
	"FybBackend/routers/v1/backend/user/token"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func SelectUserByAccount(e *gin.Engine, db *gorm.DB) {
	e.GET("/v1/backend/user/searchByAccount", func(context *gin.Context) {
		var result *multierror.Error
		mp := make(map[string]interface{})
		account := context.DefaultQuery("account", "")
		if account == "" {
			result = multierror.Append(result, errors.New("账户输入不能为空！"))
		}
		mp["account"] = account
		err1 := token.JwtVerify(context)
		user, _, err2 := fybDatabase.SelectSingleUserByCondition(db, mp)
		result = multierror.Append(result, err1, err2)
		if result.ErrorOrNil() == nil {
			context.JSON(200, gin.H{
				"code":    200,
				"message": "get userInfoList success!",
				"data":    user,
			})
		} else {
			context.JSON(404, gin.H{
				"code":    404,
				"message": result.Error(),
			})
		}
	})
}
