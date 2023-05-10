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
		var result *multierror.Error
		mp := make(map[string]interface{})

		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp)
		users, count, err3 := fybDatabase.SelectAllUserByCondition(db, mp)
		result = multierror.Append(result, err1, err2, err3)

		if result.ErrorOrNil() == nil {
			context.JSON(200, gin.H{
				"code":    200,
				"message": "get userInfoList success!",
				"data": map[string]interface{}{
					"count": count,
					"users": users,
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
