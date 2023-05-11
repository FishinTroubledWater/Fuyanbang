package selectFeedback

import (
	fybDatabase "FybBackend/database"
	"FybBackend/routers/v1/backend/token"
	"FybBackend/routers/v1/exceptionHandler"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func SelectFeedbackByPage(e *gin.Engine, db *gorm.DB) {
	e.POST("/v1/backend/feedback/list", func(context *gin.Context) {
		if err := token.JwtVerify(context); err != nil {
			context.JSON(403, gin.H{
				"code":    403,
				"message": err.Error(),
			})
			return
		}
		var result *multierror.Error
		mp := make(map[string]interface{})
		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp)
		pageNum := int64(mp["pageNum"].(float64))
		pageSize := int64(mp["pageSize"].(float64))
		posts, count, err3 := fybDatabase.SelectAllFeedbackByPage(db, "", pageNum, pageSize)
		result = multierror.Append(result, err1, err2, err3)

		code, msg := exceptionHandler.Handle(result)
		if code == 200 {
			context.JSON(code, gin.H{
				"code":    code,
				"message": "请求成功",
				"data": map[string]interface{}{
					"total":   count,
					"pageNum": pageNum,
					"posts":   posts,
				},
			})
		} else {
			context.JSON(code, gin.H{
				"code":    code,
				"message": msg,
			})
		}
	})
}
