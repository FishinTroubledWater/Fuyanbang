package login

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func PasswordLogin(e *gin.Engine, db *gorm.DB) {
	e.GET("/v1/frontend/user/passwordLogin", func(context *gin.Context) {
		var errors error
		var mp map[string]interface{}

		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp)
		user, count, err3 := fybDatabase.SelectSingleUserByCondition(db, mp)
		errors = multierror.Append(errors, err1, err2, err3)

		if count == 1 {
			context.JSON(200, gin.H{
				"code":    200,
				"message": "login success!",
				"data":    user,
			})
		} else {
			context.JSON(404, gin.H{
				"message": 404,
				"msg":     errors.Error(),
			})
		}
	})
}
