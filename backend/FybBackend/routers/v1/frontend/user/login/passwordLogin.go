package login

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

type responseItem struct {
	State string `json:"state"`
}

func PasswordLogin(e *gin.Engine, db *gorm.DB) {
	e.POST("/v1/frontend/passwordLogin", func(context *gin.Context) {
		var errors *multierror.Error
		var mp map[string]interface{}

		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp)
		_, count, err3 := fybDatabase.SelectSingleUserByCondition(db, mp)
		errors = multierror.Append(errors, err1, err2, err3)
		if count == 1 {
			context.JSON(200, gin.H{
				"code":    200,
				"message": "登陆成功",
				"data": responseItem{
					"true",
				},
			})
		} else {
			context.JSON(404, gin.H{
				"message": 404,
				"msg":     errors.Error(),
			})
		}
	})
}
