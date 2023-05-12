package myFavorites

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
	e.GET("/v1/frontend/user/myFavorites/:userID", func(context *gin.Context) {
		var result *multierror.Error
		var count int64
		var favorites []fybDatabase.FavoriteRecord
		userID := context.Param("userID")

		userIdInt64, err := strconv.ParseInt(userID, 10, 64)

		count, favorites, err = fybDatabase.SearchFavoriteByUserId(db, userIdInt64)
		if err != nil {
			result = multierror.Append(result, err)
		}

		var responseBody []map[string]interface{}
		if count > 0 {
			for i := range favorites {
				data, err := json.Marshal(&favorites[i])
				if err != nil {
					result = multierror.Append(result, err)
				}
				var postMap map[string]interface{}
				err = json.Unmarshal(data, &postMap)

				if err != nil {
					result = multierror.Append(result, err)
				}

				articleMap := postMap["Article"].(map[string]interface{})
				postMap["id"] = articleMap["ID"]
				postMap["publishTime"] = articleMap["PublishTime"]
				postMap["partID"] = articleMap["PartID"]
				postMap["title"] = articleMap["Summary"]
				postMap["favorite"] = articleMap["Favorite"]
				postMap["like"] = articleMap["Like"]
				postMap["content"] = articleMap["Content"]
				delete(postMap, "ArticleID")
				delete(postMap, "ArticleType")
				delete(postMap, "Article")
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
