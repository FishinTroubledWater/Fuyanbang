package news

import (
	fybDatabase "FybBackend/database"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func NewsInfo(e *gin.Engine, db *gorm.DB) {
	e.GET("/v1/frontend/news/detail", func(context *gin.Context) {
		var result *multierror.Error
		mp := make(map[string]interface{})
		newses, count, err1 := fybDatabase.SelectAllNewsByCondition(db, mp)
		result = multierror.Append(result, err1)

		if result.ErrorOrNil() == nil {
			context.JSON(200, gin.H{
				"code":    200,
				"message": "get all newsInfo success!",
				"data": map[string]interface{}{
					"count":  count,
					"newses": newses,
				},
			})
		} else {
			code := 404
			context.JSON(code, gin.H{
				"code":    code,
				"message": result.Error(),
			})
		}
	})
}
