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
	e.POST("/v1/frontend/academy/searchByRule", func(context *gin.Context) {
		var result error
		data, err := context.GetRawData()
		if err != nil {
			result = multierror.Append(result, err)
		}
		var m map[string]interface{}
		err = json.Unmarshal(data, &m)
		if err != nil {
			result = multierror.Append(result, err)
		}
		region := m["region"]
		level := m["level"]
		_type := m["type"]
		var responseBody []fybDatabase.Academy
		var count int64
		err, responseBody, count = fybDatabase.SearchByRegionLevelType(db, region.(string), level.(string), _type.(string))
		if err != nil {
			result = multierror.Append(result, err)
		}
		if count > 0 {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "院校筛选成功",
				"data": map[string]interface{}{
					"num":  count,
					"list": responseBody,
				},
			})
		} else if count <= 0 {
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
