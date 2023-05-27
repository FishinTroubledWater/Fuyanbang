package myFavorites

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
	"net/http"
)

func DeleteFavorite(e *gin.Engine, db *gorm.DB) {
	e.POST("/v1/frontend/user/deleteMyFavorite", func(context *gin.Context) {
		var result *multierror.Error
		var count int64
		mp := make(map[string]interface{})

		b, err1 := context.GetRawData()
		err2 := json.Unmarshal(b, &mp)

		count, err3 := fybDatabase.DeleteFavoriteByCondition(db, mp)

		result = multierror.Append(result, err1, err2, err3)
		if result.ErrorOrNil() == nil && count > 0 {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "删除成功",
			})
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": result.Error(),
			})
		}
	})
}
