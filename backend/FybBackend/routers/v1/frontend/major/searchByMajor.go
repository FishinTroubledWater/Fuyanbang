package major

import (
	fybDatabase "FybBackend/database"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
	"net/http"
)

func SearchByName(e *gin.Engine, db *gorm.DB) {
	e.GET("/v1/frontend/major/searchByName/:Name", func(context *gin.Context) {
		var result *multierror.Error
		Name := context.Param("Name")

		err, majors, count := fybDatabase.SearchMajorByName(db, Name)
		result = multierror.Append(result, err)
		if count > 0 && result.ErrorOrNil() == nil {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "专业搜索成功",
				"data": map[string]interface{}{
					"num":  count,
					"list": majors,
				},
			})
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": result.Error(),
				"data": map[string]interface{}{
					"num":  count,
					"list": majors,
				},
			})
		}
	})
}
