package newsInfo

import (
	fybDatabase "FybBackend/database"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

func NewsInfo(e *gin.Engine, db *gorm.DB) {
	e.GET("/v1/frontend/news/list", func(context *gin.Context) {
		var errors *multierror.Error
		mp := make(map[string]interface{})

		newses, _, err1 := fybDatabase.SelectAllNewsByCondition(db, mp)
		errors = multierror.Append(errors, err1)

		if errors.ErrorOrNil() == nil {
			context.JSON(200, gin.H{
				"code":    200,
				"message": "get all newsInfo success!",
				"data":    newses,
			})
		} else {
			context.JSON(404, gin.H{
				"code":    404,
				"message": errors.Error(),
			})
		}
	})
}
