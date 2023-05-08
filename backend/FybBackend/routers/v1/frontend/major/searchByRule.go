package academy

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"net/http"
)

func SearchByRule(e *gin.Engine) {
	db := fybDatabase.InitDB()
	e.POST("/v1/frontend/major/searchByRule", func(context *gin.Context) {
		var result *multierror.Error
		data, err1 := context.GetRawData()
		var mp map[string]interface{}
		err2 := json.Unmarshal(data, &mp)
		majors, _, err3 := fybDatabase.SelectMajorByCondition(db, mp)
		result = multierror.Append(result, err1, err2, err3)
		if result.ErrorOrNil() == nil {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "院校筛选成功",
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
