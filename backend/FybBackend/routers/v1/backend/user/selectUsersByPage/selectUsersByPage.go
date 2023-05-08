package selectUsersByPage

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

type responseItem struct {
	Total   int64              `json:"total"`
	PageNum int64              `json:"pageNum"`
	Users   []fybDatabase.User `json:"users"`
}

func SelectUsersByPage(e *gin.Engine, db *gorm.DB) {
	e.POST("/v1/backend/user/list", func(context *gin.Context) {
		var errors *multierror.Error
		mp := make(map[string]int64)
		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp)
		//err3 := token.JwtVerify(context)
		//pageNum := mp["pageNum"]
		//pageSize := mp["pageSize"]
		//users, count, err4 := fybDatabase.SelectAllUserByPage(db, pageNum, pageSize)
		errors = multierror.Append(errors, err1, err2)
		if errors.ErrorOrNil() == nil {
			context.JSON(200, gin.H{
				"code":    200,
				"message": "get userInfoList success!",
				//"data": responseItem{
				//	Total:   count,
				//	PageNum: pageNum,
				//	Users:   users,
				//},
			})
		} else {
			context.JSON(404, gin.H{
				"code":    404,
				"message": errors.Error(),
			})
		}
	})
}
