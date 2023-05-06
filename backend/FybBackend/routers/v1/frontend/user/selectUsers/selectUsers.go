package selectUsers

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func SelectUsers(e *gin.Engine, db *gorm.DB) {
	e.GET("/v1/frontend/user/list", func(context *gin.Context) {
		var errors *multierror.Error
		mp := make(map[string]interface{})

		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp)
		users, _, err3 := fybDatabase.SelectAllUserByCondition(db, mp)
		errors = multierror.Append(errors, err1, err2, err3)

		if errors.ErrorOrNil() == nil {
			context.JSON(200, gin.H{
				"code":    200,
				"message": "get userInfoList success!",
				"data":    users,
			})
		} else {
			context.JSON(404, gin.H{
				"code":    404,
				"message": errors.Error(),
				"data":    users,
			})
		}
	})
}
