package modifyNews

import (
	fybDatabase "FybBackend/database"
	"FybBackend/routers/v1/backend/token"
	"FybBackend/routers/v1/exceptionHandler"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func UpdateNews(e *gin.Engine, db *gorm.DB) {
	e.PATCH("/v1/backend/news/update", func(context *gin.Context) {
		if err := token.JwtVerify(context); err != nil {
			context.JSON(403, gin.H{
				"code":    403,
				"message": err.Error(),
			})
			return
		}
		var result *multierror.Error
		mp1 := make(map[string]interface{})
		mp2 := make(map[string]interface{})
		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp1)
		mp2["id"] = mp1["id"]
		delete(mp1, "id")
		_, err3 := fybDatabase.UpdateSinglePostByCondition(db, mp2, mp1)
		result = multierror.Append(result, err1, err2, err3)

		code, msg := exceptionHandler.Handle(result)
		if code == 200 {
			context.JSON(code, gin.H{
				"code":    code,
				"message": "修改成功",
			})
		} else {
			context.JSON(code, gin.H{
				"code":    code,
				"message": msg,
			})
		}
	})
}
