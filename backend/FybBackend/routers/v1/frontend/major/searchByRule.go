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
		majors, count, err3 := fybDatabase.SelectMajorByCondition(db, mp)
		result = multierror.Append(result, err1, err2, err3)
		if mp["type"] == "学科门类" {
			delete(mp, "type")
		}
		if mp["firstLevelDiscipline"] == "一级学科" {
			delete(mp, "firstLevelDiscipline")
		}
		if mp["matchType"] == "数学类型" {
			delete(mp, "matchType")
		}

		if mp["foreign"] == "外语类型" {
			delete(mp, "foreign")
		}
		if result.ErrorOrNil() == nil {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "专业筛选成功",
				"data": map[string]interface{}{
					"num":  count,
					"list": majors,
				},
			})
		} else {
			var code int
			if err3 != nil || err2 != nil {
				code = 500
			} else {
				code = 404
			}
			context.JSON(code, gin.H{
				"code":    code,
				"message": result.Error(),
			})
		}

	})
}
