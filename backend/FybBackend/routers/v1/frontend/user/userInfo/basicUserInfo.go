package userInfo

import (
	fybDatabase "FybBackend/database"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func BasicUserInfo(e *gin.Engine, db *gorm.DB) {
	e.GET("/v1/frontend/user/basicUserInfo", func(context *gin.Context) {
		var result *multierror.Error
		mp := make(map[string]interface{})
		user, _, err := fybDatabase.SelectSingleUserByCondition(db, mp)
		result = multierror.Append(result, err)
		if result.ErrorOrNil() == nil {

			context.JSON(200, gin.H{
				"code":    200,
				"message": "用户个人信息返回成功",
				"data": map[string]interface{}{
					"avatar":   user.AvatarUrl,
					"nickname": user.NickName,
					"level":    "Lv.1",
					"slogan":   user.Slogan,
					"userDays": 10,
				},
			})
		} else {
			context.JSON(404, gin.H{
				"code":    404,
				"message": result.Error(),
			})
		}
	})
}
