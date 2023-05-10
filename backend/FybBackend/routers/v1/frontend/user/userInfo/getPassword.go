package userInfo

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func GetPassword(e *gin.Engine, db *gorm.DB) {
	e.GET("/v1/frontend/user/getPassword", func(context *gin.Context) {
		var result *multierror.Error
		mp := make(map[string]interface{})

		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp)
		user, _, err3 := fybDatabase.SelectSingleUserByCondition(db, mp)
		result = multierror.Append(result, err1, err2, err3)

		if result.ErrorOrNil() == nil {
			context.JSON(200, gin.H{
				"code":    200,
				"message": "返回用户密码和邮箱成功！",
				"data": map[string]interface{}{
					"account":  user.Account,
					"password": user.Password,
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
