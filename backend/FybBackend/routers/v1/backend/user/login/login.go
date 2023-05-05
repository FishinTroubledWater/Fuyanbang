package login

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

type responseItem struct {
	Id      int    `json:"id"`
	Account string `json:"account"`
	Token   string `json:"token"`
}

func Login(e *gin.Engine, db *gorm.DB) {
	e.POST("/v1/backend/user/login", func(context *gin.Context) {
		var errors error
		mp := make(map[string]interface{})
		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp)
		user, count, err1 := fybDatabase.SelectSingleUserByCondition(db, mp)
		errors = multierror.Append(errors, err1, err2)

		if count == 1 {
			context.JSON(200, gin.H{
				"code":    200,
				"message": "frontend login success!",
				"data": responseItem{
					user.ID,
					user.Account,
					"jdioasjdiojewq",
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
