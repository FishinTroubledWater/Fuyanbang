package userInfo

import (
	fybDatabase "FybBackend/database"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

type ResponseItem struct {
	Avatar    string
	Nickname  string
	Level     string
	Slogan    string
	UsageDays int
}

func BasicUserInfo(e *gin.Engine, db *gorm.DB) {
	e.GET("/v1/frontend/user/basicUserInfo", func(context *gin.Context) {
		var errors *multierror.Error
		mp := make(map[string]interface{})
		mp["account"] = context.DefaultQuery("account", "")
		user, _, err := fybDatabase.SelectSingleUserByCondition(db, mp)
		errors = multierror.Append(errors, err)
		if errors.ErrorOrNil() == nil {
			responseItem := ResponseItem{
				Avatar:    user.AvatarUrl,
				Nickname:  user.NickName,
				Level:     "Lv.1",
				Slogan:    user.Slogan,
				UsageDays: 10,
			}
			context.JSON(200, gin.H{
				"code":    200,
				"message": "get basicUserInfo success!",
				"data":    responseItem,
			})
		} else {
			context.JSON(404, gin.H{
				"code":    404,
				"message": errors.Error(),
			})
		}
	})
}
