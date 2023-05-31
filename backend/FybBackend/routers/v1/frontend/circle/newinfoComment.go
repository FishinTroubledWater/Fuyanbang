package circle

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"net/http"
	"time"
)

func SearchNewInfoComment(e *gin.Engine) {
	db := fybDatabase.InitDB()
	e.GET("/v1/frontend/circle/newinfoComment/:postId", func(context *gin.Context) {
		var result *multierror.Error

		postId := context.Param("postId")

		err, comments := fybDatabase.SearchNewInfoComment(db, postId)
		if err != nil {
			result = multierror.Append(result, err)
		}

		var responseBody [](map[string]interface{})

		for i := range comments {
			data, err := json.Marshal(&comments[i])
			if err != nil {
				result = multierror.Append(result, err)
			}
			var commentMap map[string]interface{}
			err = json.Unmarshal(data, &commentMap)
			if err != nil {
				result = multierror.Append(result, err)
			}

			commentMap["name"] = commentMap["Author"].(map[string]interface{})["NickName"]
			if publishTime, ok := commentMap["PublishTime"].(string); ok {
				t, err := time.Parse(time.RFC3339, publishTime)
				if err == nil {
					commentMap["time"] = t.Format("2006.01.02 15:04:05")
				}
			}
			commentMap["icon"] = commentMap["Author"].(map[string]interface{})["AvatarUrl"]
			commentMap["content"] = commentMap["Content"]
			delete(commentMap, "Content")
			delete(commentMap, "Author")
			delete(commentMap, "PublishTime")
			delete(commentMap, "ID")
			delete(commentMap, "UserID")
			delete(commentMap, "TargetPost")
			delete(commentMap, "CommentNum")
			delete(commentMap, "TargetComment")

			responseBody = append(responseBody, commentMap)
		}

		if result.ErrorOrNil() == nil {
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
