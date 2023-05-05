package academy

import (
	fybDatabase "FybBackend/database"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"net/http"
)

func SelectAcademyByCode(e *gin.Engine) {
	db := fybDatabase.InitDB()
	e.GET("/v1/frontend/academy/detail/:code", func(context *gin.Context) {
		var result *multierror.Error
		code := context.Param("code")

		err, responseBody, count := fybDatabase.SearchAcademyByCode(db, code)
		result = multierror.Append(result, err)
		if count > 0 && result.ErrorOrNil() == nil {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "用户查看院校成功",
				"data":    responseBody,
			})
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": result.Error(),
				"data":    responseBody,
			})
		}
	})
}
