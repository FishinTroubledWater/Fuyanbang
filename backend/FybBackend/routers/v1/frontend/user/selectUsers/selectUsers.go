package selectUsers

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"fmt"
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
		code := 200
		if result.ErrorOrNil() == nil {
			context.JSON(code, gin.H{
				"code":    code,
				"message": "get userInfoList success!",
				"data": map[string]interface{}{
					"count": count,
					"users": users,
				},
			})
		} else {
			if err1 != nil || err2 != nil {
				code = 403
			} else if count == 0 {
				code = 404
			} else {
				code = 500
			}
			context.JSON(code, gin.H{
				"code":    code,
				"message": result.Error(),
			})
			fmt.Println(result.Error())
		}
	})
}
