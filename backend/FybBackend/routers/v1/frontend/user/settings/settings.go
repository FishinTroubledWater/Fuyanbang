package settings

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func Settings(e *gin.Engine, db *gorm.DB) {
	e.POST("/v1/frontend/user/settings", func(context *gin.Context) {
		var errors error
		mp1 := make(map[string]interface{})
		mp2 := make(map[string]interface{})

		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp1)
		account := mp1["account"]
		mp2["account"] = account
		delete(mp1, "account")

		count, err3 := fybDatabase.UpdateSingleUserByCondition(db, mp1, mp2)
		errors = multierror.Append(errors, err1, err2, err3)

		if count > 0 {
			context.JSON(200, gin.H{
				"code":    200,
				"message": "update settings success!",
				"data":    "",
			})
		} else {
			context.JSON(404, gin.H{
				"code":    404,
				"message": errors.Error(),
			})
		}
	})
}
