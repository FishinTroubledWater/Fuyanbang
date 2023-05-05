package academy

import (
	fybDatabase "FybBackend/database"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"net/http"
)

func SearchByName(e *gin.Engine) {
	db := fybDatabase.InitDB()
	e.GET("/v1/frontend/academy/searchByName/:Name", func(context *gin.Context) {
		var result error
		Name := context.Param("Name")

		err, responseBody, count := fybDatabase.SearchAcademyByName(db, Name)
		result = multierror.Append(result, err)
		if count > 0 {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "院校搜索成功",
				"data": map[string]interface{}{
					"num":  count,
					"list": responseBody,
				},
			})
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": result.Error(),
				"data": map[string]interface{}{
					"num":  count,
					"list": responseBody,
				},
			})
		}
	})
}
