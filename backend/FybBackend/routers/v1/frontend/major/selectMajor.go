package major

import (
	fybDatabase "FybBackend/database"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
	"net/http"
)

func SelectMajorByCode(e *gin.Engine, db *gorm.DB) {
	e.GET("/v1/frontend/major/detail/:code", func(context *gin.Context) {
		var result *multierror.Error
		code := context.Param("code")

		err, majors, _ := fybDatabase.SearchMajorByCode(db, code)
		result = multierror.Append(result, err)
		if result.ErrorOrNil() == nil {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "用户查看专业成功",
				"data":    majors,
			})
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": result.Error(),
			})
		}
	})
}
