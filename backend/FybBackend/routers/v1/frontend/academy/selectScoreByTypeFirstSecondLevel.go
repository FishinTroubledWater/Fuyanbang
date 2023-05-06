package academy

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"net/http"
)

func SelectScoreByTypeFirstSecondLevel(e *gin.Engine) {
	db := fybDatabase.InitDB()
	e.POST("/v1/frontend/academy/score", func(context *gin.Context) {
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
		_type := m["type"]
		firstLevelDiscipline := m["firstLevelDiscipline"]
		secondLevelDiscipline := m["secondLevelDiscipline"]

		requestBody := map[string]interface{}{
			"subjectCategory":       _type,
			"firstLevelDiscipline":  firstLevelDiscipline,
			"secondLevelDiscipline": secondLevelDiscipline,
		}

		if _type == "学科门类" {
			delete(requestBody, "subjectCategory")
		}
		if firstLevelDiscipline == "一级学科" {
			delete(requestBody, "firstLevelDiscipline")
		}
		if secondLevelDiscipline == "二级学科" {
			delete(requestBody, "secondLevelDiscipline")
		}

		var imgUrl []string

		err, imgUrl = fybDatabase.SearchScoreByTypeFirstSecondLevel(db, requestBody)
		if err != nil {
			result = multierror.Append(result, err)
		}

		if result.ErrorOrNil() == nil {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "专业分数线查看成功",
				"data": map[string]interface{}{
					"list": imgUrl,
				},
			})
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": result.Error(),
				"data": map[string]interface{}{
					"list": imgUrl,
				},
			})
		}

	})
}
