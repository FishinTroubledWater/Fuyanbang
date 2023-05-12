package circle

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"net/http"
	"strconv"
)

func SearchQueAnswer(e *gin.Engine) {
	db := fybDatabase.InitDB()
	e.GET("/v1/frontend/circle/queAnswer/:queID", func(context *gin.Context) {
		var result *multierror.Error
		var count int64
		var comments []fybDatabase.Comment
		queID := context.Param("queID")

		queIdInt64, err := strconv.ParseInt(queID, 10, 64)
		if err != nil {
			result = multierror.Append(result, err)
		}

		count, comments, err = fybDatabase.SearchQueByQueId(db, queIdInt64)
		if err != nil {
			result = multierror.Append(result, err)
		}

		var responseBody []map[string]interface{}
		if count > 0 {
			for i := range comments {
				data, err := json.Marshal(&comments[i])
				if err != nil {
					result = multierror.Append(result, err)
				}
				var postMap map[string]interface{}
				err = json.Unmarshal(data, &postMap)
				if err != nil {
					result = multierror.Append(result, err)
				}

				postMap["name"] = postMap["Author"].(map[string]interface{})["NickName"]
				postMap["time"] = postMap["PublishTime"]
				postMap["icon"] = postMap["Author"].(map[string]interface{})["AvatarUrl"]
				postMap["answer"] = postMap["Content"]
				delete(postMap, "Author")
				delete(postMap, "PublishTime")
				delete(postMap, "CommentNum")
				delete(postMap, "Content")
				delete(postMap, "TargetComment")
				delete(postMap, "TargetPost")
				delete(postMap, "UserID")
				delete(postMap, "ID")

				responseBody = append(responseBody, postMap)
			}
		}

		if result.ErrorOrNil() == nil && count > 0 {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "请求成功",
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
