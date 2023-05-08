package register

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func UserRegister(e *gin.Engine, db *gorm.DB) {
	e.POST("/v1/frontend/register", func(context *gin.Context) {
		var errors *multierror.Error
		var mp map[string]interface{}
		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp)
		count, err3 := fybDatabase.AddUser(db, mp)
		errors = multierror.Append(errors, err1, err2, err3)
		if count > 0 && errors.ErrorOrNil() == nil {
			context.JSON(200, gin.H{
				"code":    200,
				"message": "注册成功",
			})
		} else {
			context.JSON(404, gin.H{
				"message": 404,
				"msg":     errors.Error(),
			})
		}
	})
}
