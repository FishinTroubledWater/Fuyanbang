package userInfo

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func ResetPassword(e *gin.Engine, db *gorm.DB) {
	e.PATCH("/v1/frontend/user/resetPassword", func(context *gin.Context) {
		var result *multierror.Error
		mp1 := make(map[string]interface{})
		mp2 := make(map[string]interface{})
		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp1)
		mp2["password"] = mp1["password"]
		delete(mp1, "password")
		_, err3 := fybDatabase.UpdateSingleUserByCondition(db, mp1, mp2)
		result = multierror.Append(result, err1, err2, err3)
		if result.ErrorOrNil() == nil {
			context.JSON(200, gin.H{
				"code":    200,
				"message": "修改密码成功",
			})
		} else {
			context.JSON(404, gin.H{
				"message": 404,
				"msg":     result.Error(),
			})
		}
	})
}
