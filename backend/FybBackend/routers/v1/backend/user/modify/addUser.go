package modify

import (
	fybDatabase "FybBackend/database"
	"FybBackend/routers/v1/backend/user/token"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
	"time"
)

func AddUser(e *gin.Engine, db *gorm.DB) {
	e.POST("/v1/backend/user/add", func(context *gin.Context) {
		var result *multierror.Error
		mp := make(map[string]interface{})
		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp)
		mp["registerTime"] = time.Now()
		err3 := token.JwtVerify(context)
		_, err4 := fybDatabase.AddUser(db, mp)
		result = multierror.Append(result, err1, err2, err3, err4)

		code := 201
		if result.ErrorOrNil() == nil {
			context.JSON(code, gin.H{
				"code":    code,
				"message": "添加用户成功",
			})
		} else {

			if err3 != nil {
				code = 403
			} else {
				code = 500
			}
			context.JSON(code, gin.H{
				"code":    code,
				"message": result.Error(),
			})
		}
	})
}
