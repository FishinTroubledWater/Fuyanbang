package _select

import (
	fybDatabase "FybBackend/database"
	"FybBackend/routers/v1/backend/user/token"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func SelectUsersByPage(e *gin.Engine, db *gorm.DB) {
	e.POST("/v1/backend/user/list", func(context *gin.Context) {
		var result *multierror.Error
		mp := make(map[string]interface{})
		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp)
		err3 := token.JwtVerify(context)
		pageNum := int64(mp["pageNum"].(float64))
		pageSize := int64(mp["pageSize"].(float64))
		users, count, err4 := fybDatabase.SelectAllUserByPage(db, pageNum, pageSize)
		result = multierror.Append(result, err1, err2, err3, err4)
		if result.ErrorOrNil() == nil {
			context.JSON(200, gin.H{
				"code":    200,
				"message": "get userInfoList success!",
				"data": struct {
					Total   int64              `json:"total"`
					PageNum int64              `json:"pageNum"`
					Users   []fybDatabase.User `json:"users"`
				}{Total: count,
					PageNum: pageNum,
					Users:   users},
			})
		} else {
			context.JSON(404, gin.H{
				"code":    404,
				"message": result.Error(),
			})
		}
	})
}
