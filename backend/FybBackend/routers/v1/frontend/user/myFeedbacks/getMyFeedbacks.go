package myFeedbacks

import (
	fybDatabase "FybBackend/database"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"net/http"
	"strconv"
)

func SearchNewQue(e *gin.Engine) {
	db := fybDatabase.InitDB()
	e.GET("/v1/frontend/user/myFeedback/:userID", func(context *gin.Context) {
		var result *multierror.Error
		var count int64
		var feedbacks []fybDatabase.Feedback
		userID := context.Param("userID")

		userIdInt64, err := strconv.ParseInt(userID, 10, 64)

		count, feedbacks, err = fybDatabase.SearchFeedbackByUserId(db, userIdInt64)
		if err != nil {
			result = multierror.Append(result, err)
		}

		var responseBody []map[string]interface{}
		if count > 0 {
			for i := range feedbacks {
				data, err := json.Marshal(&feedbacks[i])
				if err != nil {
					result = multierror.Append(result, err)
				}
				var postMap map[string]interface{}
				err = json.Unmarshal(data, &postMap)
				if err != nil {
					result = multierror.Append(result, err)
				}

				authorMap := postMap["Author"].(map[string]interface{})
				postMap["id"] = postMap["ID"]
				postMap["content"] = postMap["Content"]
				postMap["time"] = postMap["Time"]
				postMap["state"] = postMap["State"]
				postMap["userID"] = authorMap["ID"]
				delete(postMap, "ArticleID")
				delete(postMap, "ArticleType")
				delete(postMap, "State")
				delete(postMap, "Time")
				delete(postMap, "Article")
				delete(postMap, "Content")
				delete(postMap, "Author")
				delete(postMap, "ID")
				delete(postMap, "UserID")

				responseBody = append(responseBody, postMap)
			}
		}

		if result.ErrorOrNil() == nil && count > 0 {
			context.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "获取发表过的帖子成功！",
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
