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
	User  fybDatabase.User
}

func PasswordLogin(e *gin.Engine, db *gorm.DB) {
	e.POST("/v1/frontend/passwordLogin", func(context *gin.Context) {
		var errors *multierror.Error
		var mp map[string]interface{}

		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp)
		user, _, err3 := fybDatabase.SelectSingleUserByCondition(db, mp)
		errors = multierror.Append(errors, err1, err2, err3)
		user.Password = "*********"
		if errors.ErrorOrNil() == nil {
			context.JSON(200, gin.H{
				"code":    200,
				"message": "登陆成功",
				"data": responseItem{
					"true",
					user,
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
