package userInfo

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
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
	e.GET("/v1/frontend/user/settings", func(context *gin.Context) {
		var errors error
		mp := make(map[string]interface{})

		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp)
		user, count, err3 := fybDatabase.SelectSingleUserByCondition(db, mp)
		responseItem := ResponseItem{
			Avatar:    user.AvatarUrl,
			Nickname:  user.NickName,
			Level:     "Lv.1",
			Slogan:    user.Slogan,
			UsageDays: 10,
		}
		errors = multierror.Append(errors, err1, err2, err3)

		if count > 0 {
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
