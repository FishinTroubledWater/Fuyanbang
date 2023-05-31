package recipe

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"net/http"
	"strconv"
	"time"
)

func GetRecipeList(e *gin.Engine) {
	db := fybDatabase.InitDB()
	e.GET("/v1/frontend/recipe/list/:code", func(context *gin.Context) {
		var result *multierror.Error
		var count int64
		var recipes []fybDatabase.Recipe
		code := context.Param("code")

		codeInt64, err := strconv.ParseInt(code, 10, 64)
		if err != nil {
			result = multierror.Append(result, err)
		}

		where := make(map[string]interface{})
		where["code"] = codeInt64

		recipes, count, err = fybDatabase.SelectAllRecipeByCondition(db, where)
		if err != nil {
			result = multierror.Append(result, err)
		}

		var responseBody []map[string]interface{}
		if count > 0 {
			for i := range recipes {
				data, err := json.Marshal(&recipes[i])
				if err != nil {
					result = multierror.Append(result, err)
				}
				var postMap map[string]interface{}
				err = json.Unmarshal(data, &postMap)
				if err != nil {
					result = multierror.Append(result, err)
				}
				if publishTime, ok := postMap["PublishTime"].(string); ok {
					t, err := time.Parse(time.RFC3339, publishTime)
					if err == nil {
						postMap["PublishTime"] = t.Format("2006.01.02 15:04:05")
					}
				}
				delete(postMap, "Like")
				delete(postMap, "Favorite")

				responseBody = append(responseBody, postMap)
			}
		}

		if result.ErrorOrNil() == nil && count > 0 {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "学长学姐说列表获取成功",
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
