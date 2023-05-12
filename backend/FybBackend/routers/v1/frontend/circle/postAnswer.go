package circle

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"net/http"
)

func PostAnswers(e *gin.Engine) {
	db := fybDatabase.InitDB()
	e.POST("/v1/frontend/circle/postAnswer", func(context *gin.Context) {
		var result *multierror.Error

		data, err := context.GetRawData()
		if err != nil {
			result = multierror.Append(result, err)
		}
		var m map[string]interface{}
		err = json.Unmarshal(data, &m)
		if err != nil {
			result = multierror.Append(result, err)
		}

		count, err := fybDatabase.AddComments(db, m)
		if err != nil {
			result = multierror.Append(result, err)
		}

		if result.ErrorOrNil() == nil && count > 0 {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "请求成功",
				"data":    map[string]interface{}{},
			})
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": result.Error(),
				"data":    map[string]interface{}{},
			})
		}
	})
}
