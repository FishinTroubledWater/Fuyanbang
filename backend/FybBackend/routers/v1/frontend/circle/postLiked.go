package circle

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
	"net/http"
)

func UpdatePostLiked(e *gin.Engine, db *gorm.DB) {
	e.POST("/v1/frontend/circle/postLiked", func(context *gin.Context) {
		var result *multierror.Error
		mp := make(map[string]interface{})

		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp)

		flag, err3 := fybDatabase.UpdateLikeByCondition(db, mp)

		result = multierror.Append(result, err1, err2, err3)
		if result.ErrorOrNil() == nil && flag == true {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "请求成功",
			})
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": result.Error(),
			})
		}
	})
}
